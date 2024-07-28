package db

import (
	"context"
	"errors"
	"github.com/TensoRaws/FinalRip/module/db"
	"sort"
)

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

// UpdateVideoEncodeClip 更新 Encode 后视频切片信息
func UpdateVideoEncodeClip(videoKey string, index int, encodeKey string) error {
	coll := db.DB.Collection(VIDEO_COLLECTION)

	filter := VideoClipInfo{
		Key:   videoKey,
		Index: index,
	}

	update := VideoClipInfo{
		EncodeKey: encodeKey,
	}

	_, err := coll.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}

	return nil
}

// GetVideoProgress 获取视频处理进度
func GetVideoProgress(videoKey string) (float64, error) {
	infos, err := GetVideoClips(videoKey)
	if err != nil {
		return 0, err
	}

	var progress float64
	for _, info := range infos {
		if info.EncodeKey != "" {
			progress += 1
		}
	}

	return progress / float64(len(infos)), nil
}
