package middleware

import (
	"fast_gin/utils/res"
	"github.com/gin-gonic/gin"
)

func BindJsonMiddleware[T any](c *gin.Context) {
	var cr T
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, c)
		c.Abort()
		return
	}
	c.Set("request", cr)
	return
}
func BindQueryMiddleware[T any](c *gin.Context) {
	var cr T
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		res.FailWithError(err, c)
		c.Abort()
		return
	}
	c.Set("request", cr)
	return
}
func BindUriMiddleware[T any](c *gin.Context) {
	var cr T
	err := c.ShouldBindUri(&cr)
	if err != nil {
		res.FailWithError(err, c)
		c.Abort()
		return
	}
	c.Set("request", cr)
	return
}
func GetBind[T any](c *gin.Context) T {
	// MustGet确保数据一定存在,不存在直接panic
	return c.MustGet("request").(T) // 将interface{}类型的值强制转换为泛型类型T
}
