package flags

import (
	"fast_gin/global"
	"flag"
	"fmt"
	"os"
)

type FlagOptinon struct {
	File    string
	Version bool
	DB      bool
}

var Option FlagOptinon

func Parse() {
	flag.StringVar(&Option.File, "f", "settings.yaml", "配置文件路径")
	flag.BoolVar(&Option.Version, "v", false, "打印当前版本")
	flag.BoolVar(&Option.DB, "db", false, "数据库迁移")
	flag.Parse()
	//fmt.Println(Option.File)
}

func Run() {
	if Option.DB {
		MigrateDB()
		os.Exit(0)
	}
	if Option.Version {
		fmt.Println("当前后端版本", global.Version)
		os.Exit(0)
	}
}
