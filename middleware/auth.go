package middleware

import (
	"fast_gin/utils/jwt"
	"fast_gin/utils/res"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(ctx *gin.Context) {
	token := ctx.GetHeader("token")
	_, err := jwt.CheckToken(token)
	if err != nil {
		res.FailWithMsg("认证失败", ctx)
		ctx.Abort()
	}
	ctx.Next()
}
func AdminMiddleware(ctx *gin.Context) {
	token := ctx.GetHeader("token")
	claims, err := jwt.CheckToken(token)
	if err != nil {
		res.FailWithMsg("认证失败", ctx)
		ctx.Abort()
	}
	if claims.RoleID != 1 {
		res.FailWithMsg("角色认证失败", ctx)
		ctx.Abort()
	}
	ctx.Next()
}
