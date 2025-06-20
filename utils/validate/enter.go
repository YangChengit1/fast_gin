package validate

// 错误信息显示中文
import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	"github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"strings"
)

var trans ut.Translator

func init() {
	// 创建翻译器
	uni := ut.New(zh.New())
	trans, _ = uni.GetTranslator("zh")

	// 注册翻译器
	v, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		_ = zh_translations.RegisterDefaultTranslations(v, trans)
	}
}

func ValidateError(err error) string {
	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		return err.Error() // 不是validator.ValidationErrors类型的话，就返回错误本身
	}
	// 遍历所有校验错误，翻译为中文
	var list []string
	for _, e := range errs {
		list = append(list, e.Translate(trans))
	}
	// 合并为单个字符串
	return strings.Join(list, ";")
}
