package v1

import (
	"net/http"

	"github.com/TensoRaws/FinalRip/server/internal/middleware/auth"
	"github.com/TensoRaws/FinalRip/server/internal/middleware/cros"
	"github.com/TensoRaws/FinalRip/server/internal/middleware/logger"
	"github.com/TensoRaws/FinalRip/server/internal/service/process"
	"github.com/gin-gonic/gin"
)

func NewAPI() *gin.Engine {
	r := gin.New()
	r.Use(cros.Cors())                            // 跨域中间件
	r.Use(logger.DefaultLogger(), gin.Recovery()) // 日志中间件
	r.Use(auth.RequireAuth())                     // 鉴权中间件

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "FinalRip",
		})
	})

	api := r.Group("/api/v1/")
	{
		processGroup := api.Group("process/")
		{
			// 开始压制
			processGroup.POST("start", process.Start)
			// 查看进度
			processGroup.GET("progress", process.Progress)
		}
	}

	return r
}
