package v1

import (
	"net/http"

	"github.com/TensoRaws/FinalRip/server/internal/middleware/logger"
	"github.com/TensoRaws/FinalRip/server/internal/service/process"
	"github.com/gin-gonic/gin"
)

func NewAPI() *gin.Engine {
	r := gin.New()
	r.Use(logger.DefaultLogger(), gin.Recovery()) // 日志中间件

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "FinalRip",
		})
	})

	api := r.Group("/api/v1/")
	{
		api.GET("/start", process.Start)
	}

	return r
}
