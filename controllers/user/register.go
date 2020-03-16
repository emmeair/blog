package user

import (
	"blog/utils"
	LocalValidate "blog/validate"
	"fmt"
	"github.com/gin-gonic/gin"
)

type RegisterForm struct {
	Username string `validate:"required" form:"username"`
	Email    string `validate:"required" form:"email"`
	Password string `validate:"required" form:"password"`
}

func Register(c *gin.Context) {

	var input RegisterForm

	if nil != c.BindQuery(&input) {

		utils.ResponseJson{Context: c}.ToJson(1, "参数绑定失败", nil)
		return
	}

	validate := LocalValidate.Validate{}.InitStruct(input)

	if nil != validate.Errors {

		for _, err := range validate.Errors {

			utils.ResponseJson{Context: c}.ToJson(1, err, nil)
			return
		}

	}

	fmt.Println(input)

}
