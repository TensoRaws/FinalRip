package queue

import (
	"github.com/TensoRaws/FinalRip/module/config"
	"github.com/hibiken/asynq"
	"strconv"
	"sync"
)

var (
	once sync.Once
	Qc   *asynq.Client
	Qs   *asynq.Server
)

func Init() {
	once.Do(func() {
		redisAddr := config.RedisConfig.Host + ":" + strconv.Itoa(config.RedisConfig.Port)

		// API server
		Qc = asynq.NewClient(asynq.RedisClientOpt{Addr: redisAddr, DB: 0})

		// Worker
		Qs = asynq.NewServer(
			asynq.RedisClientOpt{Addr: redisAddr, DB: 0},
			asynq.Config{Concurrency: 1},
		)
	})
}
