package pwd

import (
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

// GenerateFromPassword 加密密码，避免将真的密码放入数据库
func GenerateFromPassword(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		logrus.Errorf("加密密码错误 %s", err)
		return ""
	}
	return string(hashedPassword)
}
// CompareHashAndPassword 核验密码
func CompareHashAndPassword(hashedPassword string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
