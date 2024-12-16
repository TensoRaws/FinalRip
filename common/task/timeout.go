package task

import (
	"time"

	"github.com/hibiken/asynq"
)

// GetTaskTimeout 获取任务超时时间，当不切割时为 48 小时，切割时默认为 20 分钟，可自定义
func GetTaskTimeout(num int, timeout *int) asynq.Option {
	DEFAULT_CLIP_TIMEOUT := asynq.Timeout(20 * time.Minute)

	if num <= 1 {
		return asynq.Timeout(48 * time.Hour)
	} else {
		if timeout == nil {
			return DEFAULT_CLIP_TIMEOUT
		}
		if *timeout <= 1 {
			return DEFAULT_CLIP_TIMEOUT
		}
		return asynq.Timeout(time.Duration(*timeout) * time.Minute)
	}
}
