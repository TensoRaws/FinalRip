package db

const (
	VIDEO_COLLECTION = "video"
)

type VideoClipInfo struct {
	Key       string `bson:"key,omitempty"`
	Index     int    `bson:"index,omitempty"`
	Total     int    `bson:"total,omitempty"`
	ClipKey   string `bson:"clip_key,omitempty"`
	EncodeKey string `bson:"encode_key,omitempty"`
}
