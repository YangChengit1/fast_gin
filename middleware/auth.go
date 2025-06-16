package middleware

import (
	"fast_gin/service/redis_ser"
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
	if redis_ser.HasLogout(token) {
		res.FailWithMsg("当前登录已注销", ctx)
		ctx.Abort()
		return
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
	if redis_ser.HasLogout(token) {
		res.FailWithMsg("当前登录已注销", ctx)
		ctx.Abort()
		return
	}
	if claims.RoleID != 1 {
		res.FailWithMsg("角色认证失败", ctx)
		ctx.Abort()
	}
	ctx.Set("claims", claims)
	ctx.Next()
}
func GetAuth(c *gin.Context) (cla *jwt.MyClaims) {
	cla = new(jwt.MyClaims)
	_claims, ok := c.Get("claims")
	if !ok {
		return
	}
	cla, ok = _claims.(*jwt.MyClaims)
	return
}
