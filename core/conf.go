package core

import (
	"fmt"
	"log"
	"os"

	"github.com/Arthaslixin/FrozenThrone-go/config"
	"gopkg.in/yaml.v2"
)

// 读取yaml文件的配置
func InitConf() *config.Config {
	const ConfigFile = "settings.yaml"
	c := &config.Config{}
	yamlConf, err := os.ReadFile(ConfigFile)
	if err != nil {
		panic(fmt.Errorf("读取配置文件失败: %s", err))
	}

	err = yaml.Unmarshal(yamlConf, c)
	if err != nil {
		log.Fatalf("配置文件初始化失败: %v", err)
	}
	log.Println("配置文件初始化成功!")
	return c
}
