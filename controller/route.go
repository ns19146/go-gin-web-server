package controller

import (
	"github.com/gin-gonic/gin"
)

func StartWebServer() {
	r := gin.Default()
	r.GET("/", test)

	r.Run()
}
