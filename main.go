package main

import (
	"fast_gin/core"
	"fast_gin/flags"
	"fast_gin/global"
	"fast_gin/routers"
)

func main() {
	core.InitLogger()
	flags.Parse()
	global.Config = core.ReadConfig()
	global.DB = core.InitGorm()
	//global.Config.Db.Port = 3307 // 这里就可以进行配置的修改
	//core.DumpConfig() // 将修改的配置项进行赋值
	//logrus.Errorf("你好")
	//logrus.Infof("你好")
	global.Redis = core.InitRedis()
	flags.Run()
	routers.Run()
}
