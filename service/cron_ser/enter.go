package cron_ser

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"time"
)

func Func1() {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
}
func CronInit() {
	timezone, _ := time.LoadLocation("Asia/Shanghai")
	crontab := cron.New(cron.WithSeconds(), cron.WithLocation(timezone))
	crontab.AddFunc("* * * * * *", Func1)
	crontab.Start()
}
