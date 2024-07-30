package config

var (
	ServerConfig Server
	LogConfig    Log
	DBConfig     DB
	RedisConfig  Redis
	OSSConfig    OSS
)

func setConfig() {
	// server
	ServerConfig.Name = config.GetString("server.name")
	ServerConfig.Port = config.GetInt("server.port")
	ServerConfig.Mode = config.GetString("server.mode")
	// log
	LogConfig.Level = config.GetString("log.level")
	LogConfig.Mode = config.GetStringSlice("log.mode")
	// db
	DBConfig.Type = config.GetString("db.type")
	DBConfig.Host = config.GetString("db.host")
	DBConfig.Port = config.GetInt("db.port")
	DBConfig.Username = config.GetString("db.username")
	DBConfig.Password = config.GetString("db.password")
	DBConfig.Database = config.GetString("db.database")
	DBConfig.SSL = config.GetBool("db.ssl")
	// redis
	RedisConfig.Host = config.GetString("redis.host")
	RedisConfig.Port = config.GetInt("redis.port")
	RedisConfig.Password = config.GetString("redis.password")
	RedisConfig.PoolSize = config.GetInt("redis.poolSize")
	// oss
	OSSConfig.Type = config.GetString("oss.type")
	OSSConfig.Endpoint = config.GetString("oss.endpoint")
	OSSConfig.AccessKey = config.GetString("oss.accessKey")
	OSSConfig.SecretKey = config.GetString("oss.secretKey")
	OSSConfig.Region = config.GetString("oss.region")
	OSSConfig.Bucket = config.GetString("oss.bucket")
	OSSConfig.UseSSL = config.GetBool("oss.ssl")
	OSSConfig.HostnameImmutable = config.GetBool("oss.hostnameImmutable")
}
