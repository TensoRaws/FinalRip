package config

type Server struct {
	Name string `yaml:"name"`
	Port int    `yaml:"port"`
	Mode string `yaml:"mode"`
}

type Log struct {
	Level string   `yaml:"level"`
	Mode  []string `yaml:"mode"`
}

type DB struct {
	Type     string `yaml:"type"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
	SSL      bool   `yaml:"ssl"`
}

type Redis struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Password string `yaml:"password"`
	PoolSize int    `yaml:"poolSize"`
}

type OSS struct {
	Type              string `yaml:"type"`
	Endpoint          string `yaml:"endpoint"`
	AccessKey         string `yaml:"accessKey"`
	SecretKey         string `yaml:"secretKey"`
	Region            string `yaml:"region"`
	Bucket            string `yaml:"bucket"`
	UseSSL            bool   `yaml:"ssl"`
	HostnameImmutable bool   `yaml:"hostnameImmutable"`
}
