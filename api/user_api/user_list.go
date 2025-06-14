package user_api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (UserApi) UserListView(c *gin.Context) {
	c.String(http.StatusOK, "用户列表")
}
