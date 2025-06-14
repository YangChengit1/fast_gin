package middleware

import (
	"fast_gin/utils/jwt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthMiddleware(ctx *gin.Context) {
	token := ctx.GetHeader("token")
	_, err := jwt.CheckToken(token)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"code": 7, "msg": "认证失败", "data": gin.H{}})
		ctx.Abort()
	}
	ctx.Next()
}
func AdminMiddleware(ctx *gin.Context) {
	token := ctx.GetHeader("token")
	claims, err := jwt.CheckToken(token)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"code": 7, "msg": "认证失败", "data": gin.H{}})
		ctx.Abort()
	}
	if claims.RoleID != 1 {
		ctx.JSON(http.StatusOK, gin.H{"code": 7, "msg": "角色认证失败", "data": gin.H{}})
		ctx.Abort()
	}
	ctx.Next()
}
