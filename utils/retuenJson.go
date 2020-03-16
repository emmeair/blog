package utils

import (
	"github.com/gin-gonic/gin"
)

type ResponseJson struct {
	Context *gin.Context
}

func (r ResponseJson) ToJson(code int, msg string, data interface{}) {

	r.Context.JSON(200, map[string]interface{}{
		"code": code,
		"msg":  msg,
		"data": data,
	})
	return
}
