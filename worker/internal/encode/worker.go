package encode

import (
	"context"
	"os"
	"path"
	"strconv"
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
	mux.HandleFunc(task.VIDEO_ENCODE, Handler)

	if err := queue.Qs.Run(mux); err != nil {
		log.Logger.Fatalf("could not start worker: %v", err)
	}
}

func Handler(ctx context.Context, t *asynq.Task) error {
	var p task.EncodeTaskPayload
	if err := sonic.Unmarshal(t.Payload(), &p); err != nil {
		return err
	}
	log.Logger.Infof("Processing task ENCODE with payload %v", util.StructToString(p.Clip))

	tempSourceVideo := "source.mkv"
	tempEncodedVideo := "encoded.mkv"

	// 清理临时文件
	defer func(p ...string) {
		log.Logger.Infof("Clear temp file %v", p)
		_ = util.ClaerTempFile(p...)
	}(tempSourceVideo, tempEncodedVideo)

	err := oss.GetWithPath(p.Clip.ClipKey, tempSourceVideo)
	if err != nil {
		log.Logger.Errorf("Failed to download video %s: %v", util.StructToString(p.Clip), err)
		return err
	}

	// 等待下载完成
	log.Logger.Infof("Wait for downloading video clip %s", p.Clip.ClipKey)
	for {
		if _, err := os.Stat(tempSourceVideo); err == nil {
			break
		}
		time.Sleep(1 * time.Second)
	}
	log.Logger.Infof("Downloaded video clip %s", p.Clip.ClipKey)

	// 设置临时视频的环境变量
	err = os.Setenv("FINALRIP_SOURCE", tempSourceVideo)
	if err != nil {
		log.Logger.Errorf("Failed to set env FINALRIP_SOURCE: %v", err)
		return err
	}

	// 压制视频
	log.Logger.Infof("Start to encode video %s", util.StructToString(p.Clip))
	err = ffmpeg.EncodeVideo(p.Script, p.EncodeParam)
	if err != nil {
		log.Logger.Errorf("Failed to encode video %s: %s", util.StructToString(p.Clip), err)
		return err
	}

	// 上传压制后的视频
	key := p.Clip.Key + "-clip-encoded-" + strconv.FormatInt(int64(p.Clip.Index), 10) + path.Ext(p.Clip.Key)

	if db.CheckVideoExist(db.VideoClipInfo{
		Key:       p.Clip.Key,
		ClipKey:   p.Clip.ClipKey,
		Index:     p.Clip.Index,
		Total:     p.Clip.Total,
		EncodeKey: key,
	}) && !p.Retry {
		log.Logger.Infof("Encode Video Clip %s already exists", key)
		return nil
	}

	err = oss.PutByPath(key, tempEncodedVideo)
	if err != nil {
		log.Logger.Errorf("Failed to upload encode video %s: %s", key, err)
		return err
	}

	err = db.UpdateVideoEncodeClip(p.Clip.Key, p.Clip.ClipKey, key)
	if err != nil {
		log.Logger.Errorf("Failed to upload encode video %s: %s", key, err)
		return err
	}

	return nil
}
