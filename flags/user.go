package flags

import (
	"fast_gin/global"
	"fast_gin/models"
	"fast_gin/utils/pwd"
	"fmt"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/ssh/terminal"
	"os"
)

type User struct{}

func (User) Create() {
	var user models.UserModel
	fmt.Println("请选择角色 1、管理员 2、用户")
	_, err := fmt.Scanln(&user.RoleID)
	if err != nil {
		fmt.Println("类型输入错误", err)
		return
	}
	if user.RoleID != 1 && user.RoleID != 2 {
		fmt.Println("用户角色输入错误", err)
		return
	}
	fmt.Println("请输入用户名")
	fmt.Scanln(&user.Username)
	var u models.UserModel
	err = global.DB.Take(&u, "username = ?", user.Username).Error //查询username字段等于user.Username的一条记录，并将结果存入u变量中
	if err == nil {                                               //没有错误就证明找到了
		fmt.Println("用户名已存在")
		return
	}
	fmt.Println("请输入密码")
	//fmt.Scanln(&user.Password) // 这样输入不安全
	password, err := terminal.ReadPassword(int(os.Stdin.Fd())) // 用这个方法输入密码看不到
	if err != nil {
		fmt.Println("读取密码时出错：", err)
		return
	}
	fmt.Println("请再次输入密码")
	repassword, err := terminal.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		fmt.Println("读取密码时出错：", err)
		return
	}
	if string(password) != string(repassword) {
		fmt.Println("两次密码不一致")
		return
	}
	hashPwd := pwd.GenerateFromPassword(string(password)) // 随机产生盐值，将产生的hash密码在放入数据库
	err = global.DB.Create(&models.UserModel{
		Username: user.Username,
		Password: hashPwd,
		RoleID:   user.RoleID,
	}).Error
	if err != nil {
		logrus.Errorf("用户创建失败%s", err)
	}
	logrus.Infof("用户创建成功")
}

func (User) List() {
	var userList []models.UserModel
	global.DB.Order("create_at desc").Limit(10).Find(&userList)
	for _, model := range userList {
		fmt.Printf("用户id: %d  用户名: %s 用户昵称 %s 用户角色 %d 创建时间 %s",
			model.ID,
			model.Username,
			model.Nickname,
			model.RoleID,
			model.CreatedAt.Format("2006-01-02 15:04:05"),
		)
	}
}
