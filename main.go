package main

import (
	"fast_gin/core"
	"fast_gin/flags"
	"fast_gin/global"
)

func main() {
	flags.Parse()
	global.Config = core.ReadConfig()
	global.Config.Db.Port = 3307
}
