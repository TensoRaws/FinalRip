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
		taskGroup := api.Group("task/")
		{
			taskGroup.POST("new", task.New)
			taskGroup.POST("start", task.Start)
			taskGroup.GET("progress", task.Progress)
			taskGroup.GET("oss/presigned", task.OSSPresigned)
			taskGroup.POST("clear", task.Clear)
		}
	}

	return r
}
