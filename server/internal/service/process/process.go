package process

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/hibiken/asynq"
	"log"
	"time"
)

// 与电子邮件相关任务的有效负载。
type EmailTaskPayload struct {
	// 电子邮件接收者的ID。
	UserID int
}

func Process(c *gin.Context) {
	client := asynq.NewClient(asynq.RedisClientOpt{Addr: "localhost:6379", DB: 0})

	// 创建带有类型名称和有效负载的任务。
	payload, err := json.Marshal(EmailTaskPayload{UserID: 42})
	if err != nil {
		log.Fatal(err)
	}
	t1 := asynq.NewTask("email:welcome", payload)

	t2 := asynq.NewTask("email:reminder", payload)

	// 立即处理任务。
	info, err := client.Enqueue(t1)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf(" [*] 成功将任务加入队列：%+v", info)

	// 在24小时后处理任务。
	info, err = client.Enqueue(t2, asynq.ProcessIn(24*time.Hour))
	if err != nil {
		log.Fatal(err)
	}
	log.Printf(" [*] 成功将任务加入队列：%+v", info)
}
