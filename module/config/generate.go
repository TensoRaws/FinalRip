package config

import (
	"fmt"
)

// GenerateDSN 根据数据库类型生成相应的 DSN 字符串, 并返回数据库类型和 DSN 字符串
func GenerateDSN() (string, string, error) {
	var dsn string

	switch DBConfig.Type {
	case "mysql":
		dsn = fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
			DBConfig.Username, DBConfig.Password, DBConfig.Host, DBConfig.Port, DBConfig.Database)
	case "postgres":
		dsn = fmt.Sprintf("host=%s port=%d upload=%s password=%s dbname=%s sslmode=allow",
			DBConfig.Host, DBConfig.Port, DBConfig.Username, DBConfig.Password, DBConfig.Database)

	default: // 默认使用 mysql
		dsn = fmt.Sprintf("%v:%v@tcp(%v:%v)/%v"+
			"?charset=utf8mb4&parseTime=True&loc=Local",
			DBConfig.Username, DBConfig.Password, DBConfig.Host, DBConfig.Port, DBConfig.Database)
	}

	return DBConfig.Type, dsn, nil
}

// GenerateOSSPrefix 生成 OSS 对象存储的前缀
func GenerateOSSPrefix() string {
	var protocol string

	if OSSConfig.UseSSL {
		protocol = "https://"
	} else {
		protocol = "http://"
	}

	switch OSSConfig.Type {
	case "minio":
		OSS_PREFIX = fmt.Sprintf("%v%v/%v/", protocol, OSSConfig.Endpoint, OSSConfig.Bucket)
	case "cos":
		OSS_PREFIX = fmt.Sprintf("https://%v.%v/", OSSConfig.Bucket, OSSConfig.Endpoint)
	default:
		// 默认使用 minio
		OSS_PREFIX = fmt.Sprintf("%v%v/%v/", protocol, OSSConfig.Endpoint, OSSConfig.Bucket)
	}

	return OSS_PREFIX
}
