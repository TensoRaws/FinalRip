package db

import "time"

const (
	VIDEO_COLLECTION = "video"
	TASK_COLLECTION  = "task"
)

type VideoClipInfo struct {
	Key       string `bson:"key,omitempty"`
	Index     int    `bson:"index,omitempty"`
	Total     int    `bson:"total,omitempty"`
	ClipKey   string `bson:"clip_key,omitempty"`
	EncodeKey string `bson:"encode_key,omitempty"`
	TaskID    string `bson:"task_id,omitempty"`
}

type Task struct {
	Key         string    `bson:"key,omitempty"`
	EncodeKey   string    `bson:"encode_key,omitempty"`
	EncodeParam string    `bson:"encode_param,omitempty"`
	Script      string    `bson:"script,omitempty"`
	CreatedAt   time.Time `bson:"created_at,omitempty"`
}
