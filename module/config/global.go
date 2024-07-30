package config

import (
	"log"
)

var (
	ServerConfig Server
	LogConfig    Log
	DBConfig     DB
	RedisConfig  Redis
	OSSConfig    OSS
)

func setConfig() {
	err := config.UnmarshalKey("server", &ServerConfig)
	if err != nil {
		log.Fatalf("unable to decode into server struct, %v", err)
	}

	err = config.UnmarshalKey("log", &LogConfig)
	if err != nil {
		log.Fatalf("unable to decode into log struct, %v", err)
	}

	err = config.UnmarshalKey("db", &DBConfig)
	if err != nil {
		log.Fatalf("unable to decode into db struct, %v", err)
	}

	err = config.UnmarshalKey("redis", &RedisConfig)
	if err != nil {
		log.Fatalf("unable to decode into redis struct, %v", err)
	}

	err = config.UnmarshalKey("oss", &OSSConfig)
	if err != nil {
		log.Fatalf("unable to decode into oss struct, %v", err)
	}
}
