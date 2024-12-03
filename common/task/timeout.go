package task

import (
	"time"

	"github.com/hibiken/asynq"
)

// GetTaskTimeout 获取任务超时时间，当不切割时为 48 小时，切割时每一个clip为 1 小时
func GetTaskTimeout(num int) asynq.Option {
	if num <= 1 {
		return asynq.Timeout(48 * time.Hour)
	} else {
		return asynq.Timeout(1 * time.Hour)
	}
}
