package task

import (
	"time"

	"github.com/TensoRaws/FinalRip/common/db"
	"github.com/TensoRaws/FinalRip/module/log"
	"github.com/TensoRaws/FinalRip/module/oss"
	"github.com/TensoRaws/FinalRip/module/resp"
	"github.com/gin-gonic/gin"
)

type ProgressRequest struct {
	VideoKey string `form:"video_key" binding:"required"`
}

type ProgressResponse struct {
	EncodeKey   string `json:"encode_key"`
	EncodeParam string `json:"encode_param"`
	EncodeURL   string `json:"encode_url"`
	Progress    []bool `json:"progress"`
	Script      string `json:"script"`
}

// Progress 查看进度 (GET /progress)
func Progress(c *gin.Context) {
	// 绑定参数
	var req ProgressRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		resp.AbortWithMsg(c, err.Error())
		return
	}

	progress, err := db.GetVideoProgress(req.VideoKey)
	if err != nil {
		log.Logger.Errorf("db.GetVideoProgress failed, err: %v", err)
		resp.AbortWithMsg(c, err.Error())
		return
	}

	task, err := db.GetCompletedTask(req.VideoKey)
	if err != nil {
		log.Logger.Errorf("db.GetCompletedEncodeKey failed, err: %v", err)
		resp.AbortWithMsg(c, err.Error())
		return
	}

	var url string
	if task.EncodeKey == "" {
		log.Logger.Warnf("encode task not completed, key: %s", req.VideoKey)
		url = ""
	} else {
		url, err = oss.GetPresignedURL(task.EncodeKey, task.EncodeKey, 48*time.Hour)
		if err != nil {
			log.Logger.Errorf("oss.GetPresignedURL failed, err: %v", err)
			resp.AbortWithMsg(c, err.Error())
			return
		}
	}

	resp.OKWithData(c, &ProgressResponse{
		EncodeKey:   task.EncodeKey,
		EncodeParam: task.EncodeParam,
		EncodeURL:   url,
		Progress:    progress,
		Script:      task.Script,
	})
}
