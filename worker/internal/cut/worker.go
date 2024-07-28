package cut

import (
	"context"
	"github.com/TensoRaws/FinalRip/common/task"
	"github.com/TensoRaws/FinalRip/module/log"
	"github.com/TensoRaws/FinalRip/module/oss"
	"github.com/TensoRaws/FinalRip/module/queue"
	"github.com/bytedance/sonic"
	"github.com/hibiken/asynq"
	"path"
	"strconv"
	"time"
)

// Start starts the worker
func Start() {
	mux := asynq.NewServeMux()
	mux.HandleFunc(task.VIDEO_CUT, Handler)

	if err := queue.Qs.Run(mux); err != nil {
		log.Logger.Fatalf("could not start worker: %v", err)
	}
}

func Handler(ctx context.Context, t *asynq.Task) error {
	log.Logger.Infof("Processing task %s with payload %v", t.Type, t.Payload)
	var p task.CutTaskPayload
	if err := sonic.Unmarshal(t.Payload(), &p); err != nil {
		return err
	}

	// file format: {time}._temp.{video_format}
	tempPath := strconv.FormatInt(time.Now().Unix(), 10) + "_temp" + path.Ext(p.VideoKey)
	err := oss.GetWithPath(p.VideoKey, tempPath)
	if err != nil {
		log.Logger.Errorf("Failed to download video %s: %v", p.VideoKey, err)
		return err
	}

	return nil
}
