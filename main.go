package main

import (
	"fast_gin/core"
	"fast_gin/flags"
	"fast_gin/global"
	"github.com/sirupsen/logrus"
)

func main() {
	core.InitLogger()
	flags.Parse()
	global.Config = core.ReadConfig()
	//global.Config.Db.Port = 3307 // 这里就可以进行配置的修改
	//core.DumpConfig() // 将修改的配置项进行赋值
	logrus.Errorf("你好")
}
