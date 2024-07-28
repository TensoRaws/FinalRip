package db

const (
	VIDEO_COLLECTION = "video"
)

type VideoClipInfo struct {
	Key       string `bson:"key"`
	Index     int    `bson:"index"`
	Total     int    `bson:"total"`
	ClipKey   string `bson:"clip_key"`
	EncodeKey string `bson:"encode_key"`
}
