package common

import (
	"fast_gin/global"
	"fast_gin/models"
	"fmt"
	"gorm.io/gorm"
)

type QueryOption struct {
	models.PageInfo
	Likes    []string
	Where    *gorm.DB
	Preloads []string
	Debug    bool
}

func QueryList[T any](model T, option QueryOption) (list []T, count int64, err error) {
	list = make([]T, 0)
	query := global.DB.Where(model) // 初始化 GORM 查询链
	if option.Key != "" {           // 不为空则表示需要进行模糊匹配
		if len(option.Likes) != 0 {
			likeQuery := global.DB.Where("")
			for _, column := range option.Likes {
				likeQuery.Or(
					fmt.Sprintf("%s like ?", column), fmt.Sprintf("%%%s%%", option.Key))
			}
			query.Where(likeQuery)
		}
	}
	// 分页
	if option.Page <= 0 {
		option.Page = 1
	}
	// 预加载
	for _, preload := range option.Preloads {
		query = query.Preload(preload)
	}
	if option.Limit <= 0 {
		option.Limit = -1 // 查全部
	}
	offset := (option.Page - 1) * option.Limit
	if option.Order == "" {
		option.Order = "created_at desc"
	}
	db := global.DB.Where("")
	if option.Debug {
		db = db.Debug()
	}
	db.Where(query).Limit(option.Limit).Offset(offset).Order(option.Order).Find(&list)
	db.Model(model).Where(query).Count(&count)
	return
}
