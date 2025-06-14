package jwt

import (
	"errors"
	"fast_gin/global"
	"github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"
	"time"
)

type Claims struct {
	UserID uint `json:"userID"`
	RoleID uint `json:"roleID"`
}
type MyClaims struct {
	Claims
	jwt.RegisteredClaims
}

func SetToken(data Claims) (string, error) {
	// 实例化结构体实例
	SetClaims := MyClaims{
		Claims: data,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(global.Config.Jwt.ExpiresAt) * time.Hour)), // 过期时间
			Issuer:    global.Config.Jwt.Issuer,                                                                   // 签发人
		},
	}
	tokenStruct := jwt.NewWithClaims(jwt.SigningMethodHS256, SetClaims)
	token, err := tokenStruct.SignedString([]byte(global.Config.Jwt.Key))
	if err != nil {
		logrus.Errorf("颁发jwt失败 %s", err)
		return "", err
	}
	return token, nil
}

func CheckToken(token string) (*MyClaims, error) {
	tokenObj, err := jwt.ParseWithClaims(token, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(global.Config.Jwt.Key), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := tokenObj.Claims.(*MyClaims); ok && tokenObj.Valid {
		return claims, nil
	} else {
		return nil, errors.New("token无效")
	}
}
