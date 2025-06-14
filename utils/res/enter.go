package res

import (
	"fast_gin/utils/validate"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int
	Data any
	Msg  string
}

func Ok(data any, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code: 0,
		Data: data,
		Msg:  msg,
	})
}
func OkWithData(data any, c *gin.Context) {
	Ok(data, "成功", c)
}
func OkWithMsg(msg string, c *gin.Context) {
	Ok(gin.H{}, msg, c)
}
func Fail(code int, msg string, c *gin.Context) {
	c.JSON(http.StatusTooManyRequests, Response{
		Code: code,
		Data: gin.H{},
		Msg:  msg,
	})
}
func FailWithError(err error, c *gin.Context) {
	msg := validate.ValidateError(err)
	Fail(7, msg, c)
}
func FailWithMsg(msg string, c *gin.Context) {
	Fail(7, msg, c)
}
