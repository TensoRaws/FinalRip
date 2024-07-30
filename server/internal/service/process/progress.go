package process

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
	Progress  []bool `json:"progress"`
	EncodeUrl string `json:"encode_url"`
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

	encodeKey, err := db.GetCompletedEncodeKey(req.VideoKey)
	if err != nil {
		log.Logger.Errorf("db.GetCompletedEncodeKey failed, err: %v", err)
		resp.AbortWithMsg(c, err.Error())
		return
	}

	url, err := oss.GetPresignedURL(encodeKey, encodeKey, 48*time.Hour)
	if err != nil {
		log.Logger.Errorf("oss.GetPresignedURL failed, err: %v", err)
		resp.AbortWithMsg(c, err.Error())
		return
	}

	resp.OKWithData(c, &ProgressResponse{
		Progress:  progress,
		EncodeUrl: url,
	})
}
