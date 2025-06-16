package captcha_api

import (
	"fast_gin/utils/captcha"
	"fast_gin/utils/res"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"github.com/sirupsen/logrus"
)

type CaptchaResponse struct {
	CaptchaID string
	Captcha   string
}

func (CaptchaApi) GenerateView(c *gin.Context) {
	var driver = base64Captcha.DriverString{
		Height:          60,
		Width:           200,
		NoiseCount:      2,
		ShowLineOptions: 6,
		Length:          4,
		Source:          "0123456789",
	}
	// 使用前面定义的 driver 和 store 创建验证码对象 cp
	cp := base64Captcha.NewCaptcha(&driver, captcha.Captchastore)
	id, b64s, _, err := cp.Generate()
	if err != nil {
		logrus.Errorf("图片验证码生成失败 %s", err)
		res.FailWithMsg("图片验证码生成失败", c)
		return
	}
	res.OkWithData(CaptchaResponse{
		CaptchaID: id,
		Captcha:   b64s,
	}, c)
}
