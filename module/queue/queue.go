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
	Isp  *asynq.Inspector
)

const (
	CUT_QUEUE    = "cut_queue"
	ENCODE_QUEUE = "encode_queue"
	MERGE_QUEUE  = "merge_queue"
)

func InitServer() {
	once.Do(func() {
		redisAddr := config.RedisConfig.Host + ":" + strconv.Itoa(config.RedisConfig.Port)

		Qc = asynq.NewClient(asynq.RedisClientOpt{Addr: redisAddr, DB: 0})

		Isp = asynq.NewInspector(asynq.RedisClientOpt{Addr: redisAddr, DB: 0})
	})
}

func InitCutWorker() {
	once.Do(func() {
		redisAddr := config.RedisConfig.Host + ":" + strconv.Itoa(config.RedisConfig.Port)

		Qs = asynq.NewServer(
			asynq.RedisClientOpt{Addr: redisAddr, DB: 0},
			asynq.Config{
				Concurrency: 1,
				Queues: map[string]int{
					CUT_QUEUE: 1,
				},
			},
		)
	})
}
