package db

import (
	"context"
	"errors"
	"github.com/TensoRaws/FinalRip/module/db"
	"go.mongodb.org/mongo-driver/bson"
	"sort"
)

// CheckVideoExist 检查视频记录是否存在
func CheckVideoExist(info VideoClipInfo) bool {
	coll := db.DB.Collection(VIDEO_COLLECTION)
	count, _ := coll.CountDocuments(context.TODO(), info)
	return count > 0
}

// InsertVideo 插入视频信息
func InsertVideo(info VideoClipInfo) error {
	coll := db.DB.Collection(VIDEO_COLLECTION)
	_, err := coll.InsertOne(context.TODO(), info)
	return err
}

// InsertManyVideo 批量插入视频信息
func InsertManyVideo(infos []VideoClipInfo) error {
	coll := db.DB.Collection(VIDEO_COLLECTION)
	_, err := coll.InsertMany(context.TODO(), []interface{}{
		infos,
	})
	return err
}

// GetVideoClips 获取所有视频切片信息，按照索引排序
func GetVideoClips(videoKey string) ([]VideoClipInfo, error) {
	coll := db.DB.Collection(VIDEO_COLLECTION)

	cursor, err := coll.Find(context.TODO(), VideoClipInfo{Key: videoKey})
	if err != nil {
		return nil, err
	}

	var infos []VideoClipInfo
	if err = cursor.All(context.TODO(), &infos); err != nil {
		return nil, err
	}

	if len(infos) > 0 && len(infos) != infos[0].Total {
		return nil, errors.New("video clip not complete")
	}

	sort.Slice(infos, func(i, j int) bool {
		return infos[i].Index < infos[j].Index
	})

	return infos, nil
}

// UpdateVideo 更新视频信息
func UpdateVideo(filter VideoClipInfo, update VideoClipInfo) error {
	coll := db.DB.Collection(VIDEO_COLLECTION)

	up := bson.D{{"$set", update}}

	_, err := coll.UpdateOne(context.TODO(), filter, up)
	if err != nil {
		return err
	}
	return nil
}

// UpdateVideoEncodeClip 更新 Encode 后视频切片信息
func UpdateVideoEncodeClip(videoKey string, clipKey string, encodeKey string) error {
	coll := db.DB.Collection(VIDEO_COLLECTION)

	filter := VideoClipInfo{
		Key:     videoKey,
		ClipKey: clipKey,
	}

	update := bson.D{{"$set", VideoClipInfo{
		EncodeKey: encodeKey,
	}}}

	_, err := coll.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}

	return nil
}

// GetVideoProgress 获取视频处理进度和每个切片的状态
func GetVideoProgress(videoKey string) ([]bool, error) {
	infos, err := GetVideoClips(videoKey)
	if err != nil {
		return nil, err
	}

	var status []bool
	for _, info := range infos {
		if info.EncodeKey != "" {
			status = append(status, true)
		} else {
			status = append(status, false)
		}

	}

	return status, nil
}
