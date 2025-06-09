package main

import (
	"fast_gin/core"
	"fmt"
)

func main() {
	Cfg := core.ReadConfig()
	fmt.Println(Cfg.Db)
}
