package merge

import (
	"context"
	"github.com/TensoRaws/FinalRip/common/task"
	"github.com/TensoRaws/FinalRip/module/log"
	"github.com/TensoRaws/FinalRip/module/queue"
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
	log.Logger.Infof("Processing task Merge with payload %v", p.Clips)

	return nil
}
