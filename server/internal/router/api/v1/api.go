package v1

import (
	"net/http"

	"github.com/TensoRaws/FinalRip/server/internal/middleware/auth"
	"github.com/TensoRaws/FinalRip/server/internal/middleware/cros"
	"github.com/TensoRaws/FinalRip/server/internal/middleware/logger"
	"github.com/TensoRaws/FinalRip/server/internal/service/task"
	"github.com/gin-gonic/gin"
)

func NewAPI() *gin.Engine {
	r := gin.New()
	r.Use(cros.Cors(), logger.DefaultLogger(), gin.Recovery(), auth.RequireAuth())

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "FinalRip",
		})
	})

	api := r.Group("/api/v1/")
	{
		processGroup := api.Group("task/")
		{
			processGroup.POST("start", task.Start)
			processGroup.GET("progress", task.Progress)
			processGroup.GET("oss/presigned", task.OSSPresigned)
		}
	}

	return r
}
