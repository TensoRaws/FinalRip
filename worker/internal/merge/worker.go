package merge

import (
	"context"
	"math"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/TensoRaws/FinalRip/common/db"
	"github.com/TensoRaws/FinalRip/common/task"
	"github.com/TensoRaws/FinalRip/module/ffmpeg"
	"github.com/TensoRaws/FinalRip/module/log"
	"github.com/TensoRaws/FinalRip/module/oss"
	"github.com/TensoRaws/FinalRip/module/queue"
	"github.com/TensoRaws/FinalRip/module/util"
	"github.com/bytedance/sonic"
	"github.com/hibiken/asynq"
)

// Start starts the worker
func Start() {
	mux := asynq.NewServeMux()
	mux.HandleFunc(task.VIDEO_MERGE, Handler)

	if err := queue.Qs.Run(mux); err != nil {
		log.Logger.Fatalf("could not start worker: %v", err)
	}
}

func Handler(ctx context.Context, t *asynq.Task) error {
	var p task.MergeTaskPayload
	if err := sonic.Unmarshal(t.Payload(), &p); err != nil {
		return err
	}
	log.Logger.Infof("Processing task Merge with payload %v", p.Clips[0].Key)

	tempFolder := "temp_clips"
	tempOriginFile := "temp_source.mkv"
	tempMergedFile := "temp_merged.mkv"

	// 清理临时文件
	_ = util.ClaerTempFile(tempFolder, tempOriginFile, tempMergedFile)
	defer func(p ...string) {
		log.Logger.Infof("Clear temp file %v", p)
		_ = util.ClaerTempFile(p...)
	}(tempFolder, tempOriginFile, tempMergedFile)

	_ = os.Mkdir(tempFolder, os.ModePerm)

	var wg sync.WaitGroup
	wg.Add(1)
	// 下载原始视频
	go func() {
		defer wg.Done()

		err := oss.GetWithPath(p.Clips[0].Key, tempOriginFile)
		if err != nil {
			log.Logger.Errorf("Failed to download video %s: %v", p.Clips[0].Key, err)
			return
		}
		// 等待下载完成
		log.Logger.Infof("Downloading source video %s", p.Clips[0].Key)
		for {
			if _, err := os.Stat(tempOriginFile); err == nil {
				break
			}
			time.Sleep(1 * time.Second)
		}
		log.Logger.Infof("Downloaded source video %s", p.Clips[0].Key)
	}()

	// 下载 Encode 后的视频 Clips
	for _, clip := range p.Clips {
		wg.Add(1)
		go func(clip db.VideoClipInfo) {
			defer wg.Done()

			dlPath := tempFolder + "/" + strconv.Itoa(clip.Index) + ".mkv"

			err := oss.GetWithPath(clip.EncodeKey, dlPath)
			if err != nil {
				log.Logger.Errorf("Failed to download video clip %s: %v", clip.EncodeKey, err)
				return
			}

			// 等待下载完成
			log.Logger.Infof("Downloading encode video clip %s", clip.EncodeKey)
			for {
				if _, err := os.Stat(dlPath); err == nil {
					break
				}
				time.Sleep(1 * time.Second)
			}
			log.Logger.Infof("Downloaded encode video clip %s", clip.EncodeKey)
		}(clip)
	}

	// 等待下载完成
	wg.Wait()

	// 合并视频
	inputFiles := make([]string, len(p.Clips))
	for i := range p.Clips {
		inputFiles[i] = tempFolder + "/" + strconv.Itoa(i) + ".mkv"
	}

	err := ffmpeg.MergeVideo(tempOriginFile, inputFiles, tempMergedFile)
	if err != nil {
		log.Logger.Errorf("Failed to merge video: %v", err)
		return err
	}

	// 校验视频时长
	sourceDuration, err := ffmpeg.GetVideoDuration(tempOriginFile)
	if err != nil {
		log.Logger.Warnf("Failed to get source video duration: %v", err)
	}
	mergedDuration, err := ffmpeg.GetVideoDuration(tempMergedFile)
	if err != nil {
		log.Logger.Warnf("Failed to get merged video duration: %v", err)
	}
	if math.Abs(sourceDuration-mergedDuration) > 1 {
		log.Logger.Warnf("Merged video duration is not equal to source video duration: %v", err)
	}

	mergedKey := p.Clips[0].Key + "-Encoded" + ".mkv"
	// 上传合并后的视频
	err = oss.PutByPath(mergedKey, tempMergedFile)
	if err != nil {
		log.Logger.Errorf("Failed to upload video %s: %v", mergedKey, err)
		return err
	}

	// 保存合并后的视频信息
	err = db.UpdateUncompletedTask(p.Clips[0].Key, mergedKey)
	if err != nil {
		log.Logger.Errorf("Failed to update completed task: %v", err)
		return err
	}

	return nil
}
