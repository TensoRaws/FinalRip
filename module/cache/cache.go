package cache

import (
	"bytes"
	"context"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

var once sync.Once

type RDB uint8

var Cache *Client

type Client struct {
	C   *redis.Client
	Ctx context.Context
}

type responseWriter struct {
	gin.ResponseWriter
	b *bytes.Buffer
}

func (w responseWriter) Write(b []byte) (int, error) {
	// 向一个bytes.buffer中写一份数据来为获取body使用
	w.b.Write(b)
	// 完成gin.Context.Writer.Write()原有功能
	return w.ResponseWriter.Write(b)
}

func Init() {
	once.Do(func() {
		Cache = NewRedisClient(0)
	})
}
