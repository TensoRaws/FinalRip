package task

import (
	"os"

	"github.com/TensoRaws/FinalRip/common/db"
)

// task types, note that video: can get video:xxx:xxx format task type
// env, for self register worker
const (
	VIDEO_CUT    = "video:cut"
	VIDEO_ENCODE = "video:encode"
	VIDEO_MERGE  = "video:merge"

	FINALRIP_VIDEO_CUT_CUSTOM    = "FINALRIP_VIDEO_CUT_CUSTOM"
	FINALRIP_VIDEO_ENCODE_CUSTOM = "FINALRIP_VIDEO_ENCODE_CUSTOM"
	FINALRIP_VIDEO_MERGE_CUSTOM  = "FINALRIP_VIDEO_MERGE_CUSTOM"
)

const (
	TASK_STATUS_PENDING   = "pending"
	TASK_STATUS_RUNNING   = "running"
	TASK_STATUS_COMPLETED = "completed"
)

// CutTaskPayload is a struct that represents the payload for cut task.
type CutTaskPayload struct {
	VideoKey string `json:"video_key"`
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
}

// GetTypeName returns the custom name of the task type if it is set in the environment variables,
// typename: task.VIDEO_CUT, task.VIDEO_ENCODE, task.VIDEO_MERGE
func GetTypeName(typename string) string {
	if !CheckTaskType(typename) {
		panic("Unknown task type!!! Task Type Name: " + typename)
	}

	var env string
	switch typename {
	case VIDEO_CUT:
		env = os.Getenv(FINALRIP_VIDEO_CUT_CUSTOM)
	case VIDEO_ENCODE:
		env = os.Getenv(FINALRIP_VIDEO_ENCODE_CUSTOM)
	case VIDEO_MERGE:
		env = os.Getenv(FINALRIP_VIDEO_MERGE_CUSTOM)
	}

	if env == "" {
		return typename
	} else {
		return env
	}
}

// GetRequestTypeName returns the custom name of the task type if it is provided in the request,
// default typename: task.VIDEO_CUT, task.VIDEO_ENCODE, task.VIDEO_MERGE
func GetRequestTypeName(typename string, custom *string) string {
	if !CheckTaskType(typename) {
		panic("Unknown task type!!! Task Type Name: " + typename)
	}

	if custom == nil || *custom == "" {
		return typename
	} else {
		return *custom
	}
}

// CheckTaskType checks if the task type is valid or not
func CheckTaskType(typenames ...string) bool {
	for _, typename := range typenames {
		switch typename {
		case VIDEO_CUT, VIDEO_ENCODE, VIDEO_MERGE:
			continue
		default:
			return false
		}
	}
	return true
}
