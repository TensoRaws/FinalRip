package config

import (
	"fmt"
	"strings"
	"sync"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var (
	config *viper.Viper
	once   sync.Once
)

const ENV_PREFIX = "FINALRIP"

func Init() {
	once.Do(func() {
		initialize()
	})
}

func initialize() {
	config = viper.New()

	config.SetConfigName("finalrip")
	config.AddConfigPath("./conf/")
	config.AddConfigPath("./")
	config.AddConfigPath("$HOME/.finalrip/")
	config.AddConfigPath("/etc/finalrip/")
	config.SetConfigType("yml")

	config.SetEnvPrefix(ENV_PREFIX)
	replacer := strings.NewReplacer(".", "_")
	config.SetEnvKeyReplacer(replacer)
	config.AutomaticEnv()

	config.WatchConfig()
	config.OnConfigChange(func(e fsnotify.Event) {
		// 配置文件发生变更之后，重新初始化配置
		setConfig()
		fmt.Println("Config file changed:", e.Name)
	})

	if err := config.ReadInConfig(); err != nil {
		fmt.Println(err)
	}

	// 初始化配置
	setConfig()
	fmt.Printf("Config init success")
}

func Get(key string) interface{} {
	return config.Get(key)
}

func GetString(key string) string {
	return config.GetString(key)
}
