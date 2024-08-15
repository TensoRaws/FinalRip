package task

import (
	"github.com/TensoRaws/FinalRip/common/db"
	"github.com/TensoRaws/FinalRip/module/log"
	"github.com/TensoRaws/FinalRip/module/oss"
	"github.com/TensoRaws/FinalRip/module/resp"
	"github.com/gin-gonic/gin"
)

type ClearRequest struct {
	VideoKey string `form:"video_key" binding:"required"`
}

// Clear 清理任务 (POST /new)
func Clear(c *gin.Context) {
	// 绑定参数
	var req NewRequest
	if err := c.ShouldBind(&req); err != nil {
		resp.AbortWithMsg(c, err.Error())
		return
	}

	// 检查任务是否存在
	if !db.CheckTaskExist(req.VideoKey) {
		resp.AbortWithMsg(c, "Task not found, please check the video key")
		return
	}

	// 检查任务是否处理完成
	if !db.CheckTaskComplete(req.VideoKey) {
		// 清理任务队列
	}

	// 清理 OSS
	ossDelObjKeys := make([]string, 0)
	task, _ := db.GetTask(req.VideoKey)
	ossDelObjKeys = append(ossDelObjKeys, task.Key)
	ossDelObjKeys = append(ossDelObjKeys, task.EncodeKey)

	clips, err := db.GetVideoClips(req.VideoKey)
	if err != nil {
		resp.AbortWithMsg(c, err.Error())
		return
	}
	for _, clip := range clips {
		ossDelObjKeys = append(ossDelObjKeys, clip.ClipKey)
		ossDelObjKeys = append(ossDelObjKeys, clip.EncodeKey)
	}

	ossDelMany(ossDelObjKeys)
	log.Logger.Infof("Deleted files from OSS: %v", ossDelObjKeys)

	// 清理数据库
	err = db.DeleteTask(req.VideoKey)
	if err != nil {
		resp.AbortWithMsg(c, err.Error())
		return
	}

	err = db.DeleteVideoClips(req.VideoKey)
	if err != nil {
		resp.AbortWithMsg(c, err.Error())
		return
	}
	log.Logger.Infof("Deleted task from database: %s", req.VideoKey)

	resp.OK(c)
}

func ossDelMany(keys []string) {
	for _, key := range keys {
		err := oss.Delete(key)
		if err != nil {
			log.Logger.Errorf("Failed to delete file from OSS: %s", err)
		}
	}
}
