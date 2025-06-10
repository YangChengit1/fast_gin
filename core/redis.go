package core

import (
	"context"
	"fast_gin/global"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

func InitRedis() (client *redis.Client) {
	cfgr := global.Config.Redis
	client = redis.NewClient(&redis.Options{
		Addr:     cfgr.Addr,
		Password: cfgr.Password,
		DB:       cfgr.DB,
	})
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		logrus.Errorf("redis连接失败%s", err)
		return
	}
	logrus.Infof("redis连接成功")
	return
}
