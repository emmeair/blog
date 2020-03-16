package main

import (
	_ "blog/models"
	"blog/route"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	route.GinRouter(r)
	_ = r.Run("0.0.0.0:8080")
}
