package process

import (
	"github.com/TensoRaws/FinalRip/common/db"
	"github.com/TensoRaws/FinalRip/common/task"
	"github.com/TensoRaws/FinalRip/module/log"
	"github.com/TensoRaws/FinalRip/module/queue"
	"github.com/TensoRaws/FinalRip/module/resp"
	"github.com/TensoRaws/FinalRip/module/util"
	"github.com/bytedance/sonic"
	"github.com/gin-gonic/gin"
	"github.com/hibiken/asynq"
	"sync"
	"time"
)

type StartRequest struct {
	EncodeParam string `form:"encode_param" binding:"required"`
	Script      string `form:"script" binding:"required"`
	VideoKey    string `form:"video_key" binding:"required"`
}

// Start 开始压制 (POST /start)
func Start(c *gin.Context) {
	// 绑定参数
	var req StartRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		resp.AbortWithMsg(c, err.Error())
		return
	}

	payload, err := sonic.Marshal(task.CutTaskPayload{
		VideoKey: req.VideoKey,
	})
	if err != nil {
		log.Logger.Error("Failed to marshal payload: " + err.Error())
		resp.AbortWithMsg(c, "Failed to marshal payload")
		return
	}

	// 视频切片，上传到 OSS
	cut := asynq.NewTask(task.VIDEO_CUT, payload)

	info, err := queue.Qc.Enqueue(cut)
	if err != nil {
		log.Logger.Error("Failed to enqueue task: " + err.Error())
		resp.AbortWithMsg(c, "Failed to enqueue task")
		return
	}

	resp.AbortWithMsg(c, "Successfully enqueued task")
	// 等待任务完成

	for {
		// 检查任务是否已完成
		if info.State == asynq.TaskStateCompleted {
			break
		}

		time.Sleep(1 * time.Second)
	}

	log.Logger.Info("Cut task completed!")

	// 获取视频 clips
	clips, err := db.GetVideoClips(req.VideoKey)
	if err != nil {
		log.Logger.Error("Failed to get video clips: " + err.Error())
		return
	}

	// 等待任务完成
	var wg sync.WaitGroup
	// 开始压制任务
	for _, clip := range clips {
		payload, err := sonic.Marshal(task.EncodeTaskPayload{
			EncodeParam: req.EncodeParam,
			Script:      req.Script,
			Clip:        clip,
		})
		if err != nil {
			log.Logger.Error("Failed to marshal payload: " + err.Error())
			return
		}

		encode := asynq.NewTask(task.VIDEO_ENCODE, payload)

		info, err := queue.Qc.Enqueue(encode)
		if err != nil {
			log.Logger.Error("Failed to enqueue task: " + err.Error())
			return
		}

		log.Logger.Info("Successfully enqueued task: " + util.StructToString(clip))

		wg.Add(1)
		go func(i *asynq.TaskInfo) {
			defer wg.Done()
			// 等待任务完成
			for {
				if info.State == asynq.TaskStateCompleted {
					break
				}

				time.Sleep(1 * time.Second)
			}
		}(info)
	}

	// 等待所有任务完成
	wg.Wait()
	log.Logger.Info("All Encode tasks completed!")

	// 开始合并任务
	clips, err = db.GetVideoClips(req.VideoKey)
	if err != nil {
		log.Logger.Error("Failed to get video clips: " + err.Error())
		return
	}
	payload, err = sonic.Marshal(task.MergeTaskPayload{
		Clips: clips,
	})
	if err != nil {
		log.Logger.Error("Failed to marshal payload: " + err.Error())
		return
	}

	merge := asynq.NewTask(task.VIDEO_MERGE, payload)

	info, err = queue.Qc.Enqueue(merge)
	if err != nil {
		log.Logger.Error("Failed to enqueue task: " + err.Error())
		return
	}

	log.Logger.Info("Successfully enqueued task: merge")

	// 等待任务完成
	for {
		if info.State == asynq.TaskStateCompleted {
			break
		}

		time.Sleep(1 * time.Second)
	}

	log.Logger.Info("Merge task completed!")
}
