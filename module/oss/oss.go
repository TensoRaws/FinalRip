package oss

import (
	"bytes"
	"context"
	"io"
	"net/url"
	"sync"
	"time"

	"github.com/TensoRaws/FinalRip/module/config"
	"github.com/TensoRaws/FinalRip/module/log"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var (
	oss  *minio.Client
	err  error
	once sync.Once
)

func Init() {
	once.Do(func() {
		initialize()
	})
}

func initialize() {
	// Initialize minio client object.
	oss, err = minio.New(config.OSSConfig.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(config.OSSConfig.AccessKey, config.OSSConfig.SecretKey, ""),
		Secure: config.OSSConfig.UseSSL,
	})
	err = oss.MakeBucket(context.TODO(), config.OSSConfig.Bucket,
		minio.MakeBucketOptions{Region: config.OSSConfig.Region, ObjectLocking: false})
	if err != nil {
		exists, _ := oss.BucketExists(context.TODO(), config.OSSConfig.Bucket)
		if !exists {
			log.Logger.Error("Failed to create bucket: " + err.Error())
		}
	}
}

// Put saves the content read from r to the oss key.
func Put(key string, reader io.Reader, objectSize int64) error {
	info, err := oss.PutObject(context.Background(), config.OSSConfig.Bucket, key, reader, objectSize,
		minio.PutObjectOptions{ContentType: "application/octet-stream"})

	log.Logger.Infof("Uploaded %s of size: %v Successfully.", key, info.Size)
	return err
}

// PutBytes saves the byte array to the oss key.
func PutBytes(key string, data []byte) error {
	return Put(key, bytes.NewReader(data), int64(len(data)))
}

// Get gets the file pointed to by key.
func Get(key string) (*minio.Object, error) {
	return oss.GetObject(context.Background(), config.OSSConfig.Bucket, key, minio.GetObjectOptions{})
}

// GetBytes gets the file pointed to by key and returns a byte array.
func GetBytes(key string) ([]byte, error) {
	obj, err := Get(key)
	if err != nil {
		return nil, err
	}
	defer func(obj *minio.Object) {
		err := obj.Close()
		if err != nil {
			log.Logger.Error("Failed to close object: " + err.Error())
		}
	}(obj)

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(obj)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// GetPresignedURL gets the presigned URL for the file pointed to by key.
func GetPresignedURL(key string, fileName string, expiration time.Duration) (string, error) {
	// Set request parameters
	reqParams := make(url.Values)
	reqParams.Set("response-content-disposition", "attachment; filename="+fileName)

	// Generate presigned get object url
	presignedURL, err := oss.PresignedGetObject(context.Background(), config.OSSConfig.Bucket, key, expiration, reqParams)
	if err != nil {
		log.Logger.Error("Failed to generate presigned URL: " + key + err.Error())
		return "", err
	}
	return presignedURL.String(), nil
}
