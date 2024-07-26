package encode

import (
	"context"
	"encoding/json"
	"github.com/hibiken/asynq"
	"log"
)

// 与电子邮件相关任务的有效负载。
type EmailTaskPayload struct {
	// 电子邮件接收者的ID。
	UserID int
}

func Start() {
	srv := asynq.NewServer(
		asynq.RedisClientOpt{Addr: "localhost:6379"},
		asynq.Config{Concurrency: 10},
	)

	mux := asynq.NewServeMux()
	mux.HandleFunc("email:welcome", sendWelcomeEmail)
	mux.HandleFunc("email:reminder", sendReminderEmail)

	if err := srv.Run(mux); err != nil {
		log.Fatal(err)
	}
}

func sendWelcomeEmail(ctx context.Context, t *asynq.Task) error {
	var p EmailTaskPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return err
	}
	log.Printf(" [*] 给用户 %d 发送欢迎邮件", p.UserID)
	return nil
}

func sendReminderEmail(ctx context.Context, t *asynq.Task) error {
	var p EmailTaskPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return err
	}
	log.Printf(" [*] 给用户 %d 发送提醒邮件", p.UserID)
	return nil
}
