package task

const (
	VIDEO_CUT    = "video:cut"
	VIDEO_ENCODE = "video:encode"
	VIDEO_MERGE  = "video:merge"
)

// CutTaskPayload is a struct that represents the payload for cut task.
type CutTaskPayload struct {
	VideoKey string `json:"video_key"`
}

// EncodeTaskPayload is a struct that represents the payload for encode task.
type EncodeTaskPayload struct {
	EncodeParam string `json:"encode_param"`
	Script      string `json:"script"`
	ClipKey     string `json:"clip_key"`
}

// MergeTaskPayload is a struct that represents the payload for merge task.
type MergeTaskPayload struct {
	VideoKey string   `json:"video_key"`
	ClipKeys []string `json:"clip_keys"`
}
