package encode

import (
	"context"
	"encoding/json"
	"log"

	"github.com/TensoRaws/FinalRip/module/queue"
	"github.com/hibiken/asynq"
)

// 与电子邮件相关任务的有效负载。
type EmailTaskPayload struct {
	// 电子邮件接收者的ID。
	UserID int
}

func Start() {
	mux := asynq.NewServeMux()
	mux.HandleFunc("email:welcome", sendWelcomeEmail)
	mux.HandleFunc("email:reminder", sendReminderEmail)

	if err := queue.Qs.Run(mux); err != nil {
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
