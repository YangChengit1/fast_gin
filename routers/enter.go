package routers

import (
	"fast_gin/global"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Run() {
	gin.SetMode(global.Config.System.Mode)
	r := gin.Default()
	r.Static("/uploads", "uploads")
	g := r.Group("api")
	UserRouter(g)
	addr := global.Config.System.Addr()
	if global.Config.System.Mode == "release" {
		logrus.Infof("后端服务运行在 %s", addr)
	}
	r.Run(addr) // Addr() 是一个方法，而不是 函数（Function），所以必须通过 结构体实例global.Config来调用
}
