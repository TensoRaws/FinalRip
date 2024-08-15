package task

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
)

type RetryEncodeRequest struct {
	VideoKey    string `form:"video_key" binding:"required"`
	EncodeParam string `form:"encode_param" binding:"required"`
	Index       *int   `form:"index" binding:"required"`
	Script      string `form:"script" binding:"required"`
}

// RetryEncode 重试压制任务 (POST /retry/encode)
func RetryEncode(c *gin.Context) {
	// 绑定参数
	var req RetryEncodeRequest
	if err := c.ShouldBind(&req); err != nil {
		resp.AbortWithMsg(c, err.Error())
		return
	}

	info := db.VideoClipInfo{
		Key:     req.VideoKey,
		ClipKey: util.GenerateClipKey(req.VideoKey, *req.Index),
	}

	// 检查任务是否存在
	if !db.CheckVideoExist(info) {
		log.Logger.Error("Encode clip not exist: " + util.StructToString(info))
		resp.AbortWithMsg(c, "Encode clip not exist")
		return
	}

	clip, err := db.GetVideoClip(info)
	if err != nil {
		log.Logger.Error("Get video clip info failed: " + err.Error())
		resp.AbortWithMsg(c, "Get video clip info failed: "+err.Error())
		return
	}

	// 取消先前的任务
	CancelTask(clip.TaskID)

	resp.OK(c)
	// 重新创建任务
	go HandleRetryEncode(req, clip)
}

func HandleRetryEncode(req RetryEncodeRequest, clip db.VideoClipInfo) {
	payload, err := sonic.Marshal(task.EncodeTaskPayload{
		EncodeParam: req.EncodeParam,
		Script:      req.Script,
		Clip:        clip,
		Retry:       true,
	})
	if err != nil {
		log.Logger.Error("Failed to marshal payload: " + err.Error())
		return
	}

	encode := asynq.NewTask(task.VIDEO_ENCODE, payload)

	info, err := queue.Qc.Enqueue(encode, asynq.Queue(queue.ENCODE_QUEUE))
	if err != nil {
		log.Logger.Error("Failed to Re-enqueue task: " + err.Error())
		return
	}

	err = db.UpdateVideo(db.VideoClipInfo{Key: req.VideoKey, ClipKey: clip.ClipKey},
		db.VideoClipInfo{TaskID: info.ID})
	if err != nil {
		log.Logger.Error("Failed to Re-enqueue task: " + err.Error())
		return
	}

	log.Logger.Info("Successfully Re-enqueued task: " + util.StructToString(clip))
}
