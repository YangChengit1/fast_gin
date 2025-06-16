package redis_ser

import (
	"context"
	"fast_gin/global"
	"fast_gin/utils/jwt"
	"fmt"
	"github.com/sirupsen/logrus"
	"time"
)

func Logout(token string) {
	claims, err := jwt.CheckToken("token")
	if err != nil {
		return
	}
	key := fmt.Sprintf("logout_%s", token)                                 // 构造Redis中的键
	sub := claims.ExpiresAt.Sub(time.Now())                                // sub表示距离过期还剩多长时间
	_, err = global.Redis.Set(context.Background(), key, "", sub).Result() // 使用Set方法将键写入Redis，值为""，有效期为sub
	if err != nil {
		logrus.Error(err)
	}
}

func HasLogout(token string) bool {
	key := fmt.Sprintf("logout_%s", token)
	_, err := global.Redis.Get(context.Background(), key).Result() // 判断key值是否在redis中
	if err == nil {
		// 在里面
		return true
	}
	return false
}
