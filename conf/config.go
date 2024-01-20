package conf

import (
	"fmt"
	yaml "gopkg.in/yaml.v3"
	"io/ioutil"
	"sync"
)

type Config struct {
	Authorization string   `yaml:"authorization"`
	XGuestToken   string   `yaml:"x-guest-token"`
	UserIDs       []string `yaml:"user-ids"`
	Cookie        string   `yaml:"cookie"`
	XCSRFToken    string   `yaml:"x-csrf-token"`
}

var (
	configOnce    sync.Once
	defaultConfig Config
)

func GetConf() Config {
	return defaultConfig
}

func InitConfig(configPath string) {
	configOnce.Do(func() {
		fmt.Printf("初始化配置文件，文件地址: %v\n", configPath)
		file, err := ioutil.ReadFile(configPath)
		if err != nil {
			panic(err)
		}

		err = yaml.Unmarshal(file, &defaultConfig)
		if err != nil {
			panic(err)
		}
		fmt.Println("配置文件初始化完成")
	})
}