package queue

import (
	"strconv"
	"sync"

	"github.com/TensoRaws/FinalRip/common/version"
	"github.com/TensoRaws/FinalRip/module/config"
	"github.com/hibiken/asynq"
)

var (
	once sync.Once
	Qc   *asynq.Client
	Qs   *asynq.Server
	Isp  *asynq.Inspector
)

const (
	CUT_QUEUE    = "cut_" + version.FINALRUP_VERSION
	ENCODE_QUEUE = "encode_" + version.FINALRUP_VERSION
	MERGE_QUEUE  = "merge_" + version.FINALRUP_VERSION
)

func getRedisClientOpt() asynq.RedisClientOpt {
	return asynq.RedisClientOpt{
		Addr:     config.RedisConfig.Host + ":" + strconv.Itoa(config.RedisConfig.Port),
		Password: config.RedisConfig.Password,
		PoolSize: config.RedisConfig.PoolSize,
		DB:       0,
	}
}

func InitServer() {
	once.Do(func() {
		Qc = asynq.NewClient(getRedisClientOpt())

		Isp = asynq.NewInspector(getRedisClientOpt())
	})
}

func InitCutWorker() {
	once.Do(func() {
		Qs = asynq.NewServer(
			getRedisClientOpt(),
			asynq.Config{
				Concurrency: 1,
				Queues: map[string]int{
					CUT_QUEUE: 1,
				},
			},
		)
	})
}

func InitEncodeWorker() {
	once.Do(func() {
		Qs = asynq.NewServer(
			getRedisClientOpt(),
			asynq.Config{
				Concurrency: 1,
				Queues: map[string]int{
					ENCODE_QUEUE: 1,
				},
			},
		)
	})
}

func InitMergeWorker() {
	once.Do(func() {
		Qs = asynq.NewServer(
			getRedisClientOpt(),
			asynq.Config{
				Concurrency: 1,
				Queues: map[string]int{
					MERGE_QUEUE: 1,
				},
			},
		)
	})
}
