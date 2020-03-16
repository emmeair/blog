package route

import (
	"blog/controllers/user"
	"github.com/gin-gonic/gin"
)

func GinRouter(c *gin.Engine) *gin.Engine {

	v1 := c.Group("/v1")
	{
		api := v1.Group("/api")

		api.GET("/login", user.Login)
		api.POST("/register", user.Register)

	}

	return c
}
