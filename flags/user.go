package flags

import (
	"fast_gin/models"
	"fmt"
	"golang.org/x/crypto/ssh/terminal"
	"os"
)

type User struct {
}

func (User) Create() {
	var user models.UserModel
	fmt.Println("请选择角色 1、管理员 2、用户")
	_, err := fmt.Scanln(&user.RoleID)
	if err != nil {
		fmt.Println("输入错误", err)
		return
	}
	if user.RoleID != 1 && user.RoleID != 2 {
		fmt.Println("用户角色输入错误", err)
		return
	}
	fmt.Println("请输入用户名")
	fmt.Scanln(&user.Username)
	fmt.Println("请输入密码")
	//fmt.Scanln(&user.Password) // 这样输入不安全
	password, err := terminal.ReadPassword(int(os.Stdin.Fd()))
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
	user.Password = string(password)
	fmt.Println(user.RoleID)
	fmt.Println(user.Username)
	fmt.Println(user.Password)
}
func (User) List() {

}
