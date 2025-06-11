package api

import "fast_gin/api/user_api"

type Api struct {
	UserApi user_api.UserApi
	name    string
}

var App = new(Api) // 指针类型
