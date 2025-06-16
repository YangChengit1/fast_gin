package models

import "gorm.io/gorm"

type UserModel struct {
	gorm.Model
	Username string `gorm:"size:16" json:"username"`
	Nickname string `gorm:"size:32" json:"nickname"`
	Password string `gorm:"size:64" json:"-"` // 不能把密码露出来
	RoleID   int8   `json:"roleID"`
}
