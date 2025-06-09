package flags

import (
	"flag"
	"fmt"
)

type FlagOptinon struct {
	File string
}

var Option FlagOptinon

func Parse() {
	flag.StringVar(&Option.File, "f", "settings.yaml", "配置文件路径")
	flag.Parse()
	fmt.Println(Option.File)
}
