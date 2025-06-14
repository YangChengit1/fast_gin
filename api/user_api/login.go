package user_api

import (
	"fast_gin/middleware"
	"fast_gin/utils/res"
	"fmt"
	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Username string `json:"username" binding:"required"` //必须提供 username
	Password string `json:"password" binding:"required"` //必须提供 password
}

func (UserApi) LoginView(c *gin.Context) {
	cr := middleware.GetBind[LoginRequest](c)
	fmt.Println(cr)
	res.OkWithData("用户登录", c)
	return
}
