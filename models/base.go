package models

type PageInfo struct {
	Page  int    `form:"page"`
	Limit int    `form:"limit"`
	Key   string `form:"key"` // 模糊匹配
	Order string `form:"order"`
}
