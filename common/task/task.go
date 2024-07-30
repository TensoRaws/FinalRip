package task

import "github.com/TensoRaws/FinalRip/common/db"

const (
	VIDEO_CUT    = "video:cut"
	VIDEO_ENCODE = "video:encode"
	VIDEO_MERGE  = "video:merge"
)

// CutTaskPayload is a struct that represents the payload for cut task.
type CutTaskPayload struct {
	VideoKey string `json:"video_key"`
	Retry    bool   `json:"retry"`
}

// EncodeTaskPayload is a struct that represents the payload for encode task.
type EncodeTaskPayload struct {
	EncodeParam string           `json:"encode_param"`
	Script      string           `json:"script"`
	Clip        db.VideoClipInfo `json:"clip"`
	Retry       bool             `json:"retry"`
}

// MergeTaskPayload is a struct that represents the payload for merge task.
type MergeTaskPayload struct {
	Clips []db.VideoClipInfo `json:"clips"`
	Retry bool               `json:"retry"`
}
