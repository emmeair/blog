package user

import (
	"blog/models"
	"blog/models/user"
	"blog/utils"
	LocalValidate "blog/validate"
	"github.com/gin-gonic/gin"
)

type RegisterForm struct {
	Username        string `validate:"required" form:"username"`
	Email           string `validate:"required" form:"email"`
	Password        string `validate:"eqfield=ConfirmPassword" form:"password"`
	ConfirmPassword string `validate:"required" form:"confirm_password"`
}

func Register(c *gin.Context) {

	var input RegisterForm

	if nil != c.Bind(&input) {

		utils.ResponseJson{Context: c}.ToJson(-1, "参数绑定失败", nil)
		return
	}

	validate := LocalValidate.Validate{}.InitStruct(input)

	if nil != validate.Errors {

		for _, err := range validate.Errors {

			utils.ResponseJson{Context: c}.ToJson(-1, err, nil)
			return
		}

	}
	checkMember := user.Members{}

	models.Db.Where("email = ?", input.Email).First(&checkMember)

	if checkMember.ID != 0 {

		utils.ResponseJson{Context: c}.ToJson(-1, "该邮箱已经被使用", nil)
		return
	}
	member := user.Members{}

	member.Username = input.Username
	member.Email = input.Email
	member.Password = input.Password
	tx := models.Db.Begin()

	if tx.Create(&member).Error != nil {
		tx.Rollback()
		utils.ResponseJson{Context: c}.ToJson(-1, "创建账户失败", nil)
		return
	}

	tx.Commit()
	utils.ResponseJson{Context: c}.ToJson(1, "创建账户成功", nil)
	return
}
