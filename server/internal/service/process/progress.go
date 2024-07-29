package process

import (
	"github.com/TensoRaws/FinalRip/common/db"
	"github.com/TensoRaws/FinalRip/module/resp"
	"github.com/gin-gonic/gin"
)

type ProgressRequest struct {
	VideoKey string `form:"video_key" binding:"required"`
}

type ProgressResponse struct {
	Progress []bool `json:"progress"`
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
		resp.AbortWithMsg(c, err.Error())
		return
	}

	resp.OKWithData(c, &ProgressResponse{
		Progress: progress,
	})
}
