package LocalValidate

import (
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh2 "github.com/go-playground/validator/v10/translations/zh"
)

type Validate struct {
	trans  ut.Translator
	Errors []string
}

func (l Validate) InitStruct(i interface{}) Validate {

	//调整语言
	l.zhTrans()

	validate := validator.New()
	_ = zh2.RegisterDefaultTranslations(validate, l.trans)
	err := validate.Struct(i)
	var validatorSile []string

	if nil != err {
		for _, err := range err.(validator.ValidationErrors) {

			validatorSile = append(validatorSile, err.Translate(l.trans))
		}
	}

	l.Errors = validatorSile

	return l
}
func (l *Validate) zhTrans() {
	uni := ut.New(zh.New())
	trans, _ := uni.GetTranslator("zh")
	l.trans = trans
}
