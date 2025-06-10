package core

import (
	"fast_gin/config"
	"fast_gin/flags"
	"fast_gin/global"
	"fmt"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
	"os"
)

func ReadConfig() (cfg *config.Config) {
	cfg = new(config.Config)
	byteData, err := os.ReadFile("settings.yaml") // 读取程序运行时的目录(fast_gin) 下的 settings.yaml 文件
	if err != nil {
		logrus.Fatalf("配置文件读取错误 %s", err)
		return
	}
	err = yaml.Unmarshal(byteData, cfg)
	if err != nil {
		logrus.Fatalf("配置文件格式错误 %s", err)
		return
	}
	logrus.Infof("%s 配置文件读取成功", flags.Option.File)
	return
}
func DumpConfig() {
	bytedata, err := yaml.Marshal(global.Config)
	if err != nil {
		fmt.Printf("配置文件转换错误 %s", err)
		return
	}
	err = os.WriteFile(flags.Option.File, bytedata, 0666)
	if err != nil {
		logrus.Errorf("配置文件写入错误 %s", err)
		return
	}
	logrus.Infof("配置文件写入成功")
}
