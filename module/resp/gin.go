package resp

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// OK 返回成功信息
func OK(c *gin.Context) {
	resp := &gin.H{
		"success": true,
	}
	c.JSON(http.StatusOK, resp)
}

// OKWithData 返回成功信息，携带自定义数据（结构体）
func OKWithData(c *gin.Context, data interface{}) {
	resp := &gin.H{
		"success": true,
		"data":    data,
	}

	c.JSON(http.StatusOK, resp)
}

// AbortWithMsg 返回错误码，自定义错误信息
func AbortWithMsg(c *gin.Context, msg string) {
	errorResp := &gin.H{
		"message": msg,
	}
	resp := &gin.H{
		"success": false,
		"error":   &errorResp,
	}
	c.AbortWithStatusJSON(http.StatusOK, resp)
}
