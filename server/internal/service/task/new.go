package task

import (
	"github.com/TensoRaws/FinalRip/common/db"
	"github.com/TensoRaws/FinalRip/module/log"
	"github.com/TensoRaws/FinalRip/module/resp"
	"github.com/gin-gonic/gin"
)

type NewRequest struct {
	VideoKey string `form:"video_key" binding:"required"`
}

// New 创建任务 (POST /new)
func New(c *gin.Context) {
	// 绑定参数
	var req NewRequest
	if err := c.ShouldBind(&req); err != nil {
		resp.AbortWithMsg(c, err.Error())
		return
	}

	// 检查任务是否存在
	if db.CheckTaskExist(req.VideoKey) {
		resp.AbortWithMsg(c, "Task already exists, please wait for it to complete or delete it.")
		return
	}

	err := db.InsertTask(req.VideoKey)
	if err != nil {
		log.Logger.Error("Failed to insert task: " + err.Error())
		resp.AbortWithMsg(c, err.Error())
		return
	}

	resp.OK(c)
}
