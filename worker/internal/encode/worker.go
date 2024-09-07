package encode

import (
	"context"
	"errors"
	"os"
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
	_ = util.ClaerTempFile(tempSourceVideo, tempEncodedVideo)
	defer func(p ...string) {
		log.Logger.Infof("Clear temp file %v", p)
		_ = util.ClaerTempFile(p...)
	}(tempSourceVideo, tempEncodedVideo)

	// 等待下载完成
	log.Logger.Infof("Wait for downloading video clip %s", p.Clip.ClipKey)

	err := oss.GetWithPath(p.Clip.ClipKey, tempSourceVideo)
	if err != nil {
		log.Logger.Errorf("Failed to download video %s: %v", util.StructToString(p.Clip), err)
		return err
	}
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

	// 检查文件大小
	if util.GetFileSize(tempEncodedVideo) < 8192 {
		log.Logger.Errorf("Failed to encode video %s: file size is too small", util.StructToString(p.Clip))
		return errors.New("file size is too small")
	}

	key := util.GenerateClipEncodedKey(p.Clip.Key, p.Clip.Index)

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

	// 检查任务是否被取消
	if !db.CheckVideoExist(db.VideoClipInfo{
		Key:     p.Clip.Key,
		ClipKey: p.Clip.ClipKey,
	}) {
		log.Logger.Errorf("Encode Video Clip %s has been canceled", key)
		return errors.New("encode video clip has been canceled")
	}

	// 上传压制后的视频
	err = oss.PutByPath(key, tempEncodedVideo)
	if err != nil {
		log.Logger.Errorf("Failed to upload encode video %s: %s", key, err)
		return err
	}

	err = db.UpdateVideoClip(db.VideoClipInfo{Key: p.Clip.Key, ClipKey: p.Clip.ClipKey}, db.VideoClipInfo{EncodeKey: key})
	if err != nil {
		log.Logger.Errorf("Failed to update video clip %s: %s", key, err)
		return err
	}

	return nil
}
