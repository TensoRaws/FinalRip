package config

import (
	"fmt"
	"os"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/TensoRaws/FinalRip/module/util"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
)

var (
	config *viper.Viper
	once   sync.Once
)

const (
	ENV_PREFIX                  = "FINALRIP"
	FINALRIP_REMOTE_CONFIG_HOST = "FINALRIP_REMOTE_CONFIG_HOST"
	FINALRIP_REMOTE_CONFIG_KEY  = "FINALRIP_REMOTE_CONFIG_KEY"
)

func Init() {
	once.Do(func() {
		initialize()
	})
}

func initialize() {
	config = viper.New()
	config.SetConfigType("yml")
	config.SetConfigName("finalrip")

	if os.Getenv(FINALRIP_REMOTE_CONFIG_HOST) == "" {
		config.AddConfigPath("./conf/")
		config.AddConfigPath("./")
		config.AddConfigPath("$HOME/.finalrip/")
		config.AddConfigPath("/etc/finalrip/")
		if err := config.ReadInConfig(); err != nil {
			fmt.Println(err)
		}
		config.WatchConfig()
	} else {
		// 从 consul 读取配置
		host := os.Getenv(FINALRIP_REMOTE_CONFIG_HOST)
		key := os.Getenv(FINALRIP_REMOTE_CONFIG_KEY)
		if key == "" {
			fmt.Println("remote key is empty, default to finalrip.yml")
			key = "finalrip.yml"
		}

		err := config.AddRemoteProvider("consul", host, key)
		if err != nil {
			fmt.Println("failed to add remote provider: " + err.Error())
		}
		err = config.ReadRemoteConfig()
		if err != nil {
			fmt.Println("failed to read remote config: " + err.Error())
		}
		go func() {
			for {
				updateRemoteConfigOnChange()
				time.Sleep(5 * time.Second)
			}
		}()
	}

	config.SetEnvPrefix(ENV_PREFIX)
	replacer := strings.NewReplacer(".", "_")
	config.SetEnvKeyReplacer(replacer)
	config.AutomaticEnv()

	config.OnConfigChange(func(e fsnotify.Event) {
		// 配置文件发生变更之后，重新初始化配置
		setConfig()
		fmt.Println("Config file changed:", e.Name)
	})

	// 初始化配置
	setConfig()
}

func updateRemoteConfigOnChange() {
	lastCfg, err := util.DeepCopyMap(config.AllSettings())
	if err != nil {
		fmt.Println("failed to copy config: " + err.Error())
	}

	err = config.WatchRemoteConfig()
	if err != nil {
		fmt.Println("failed to watch remote config: " + err.Error())
	}

	cfg, err := util.DeepCopyMap(config.AllSettings())
	if err != nil {
		fmt.Println("failed to copy config: " + err.Error())
	}

	if !reflect.DeepEqual(lastCfg, cfg) {
		// 配置文件发生变更之后，重新初始化配置
		setConfig()
		fmt.Println("Remote Config file changed!")
	}
}
