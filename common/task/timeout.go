package task

import (
	"time"

	"github.com/hibiken/asynq"
)

// GetTaskTimeout 获取任务超时时间，当不切割时为 48 小时，切割时默认为 20 分钟，可自定义
func GetTaskTimeout(num int, deadline *int) asynq.Option {
	if num <= 1 {
		return asynq.Timeout(48 * time.Hour)
	} else {
		if deadline == nil {
			return asynq.Timeout(20 * time.Minute)
		}
		return asynq.Timeout(time.Duration(*deadline) * time.Minute)
	}
}
