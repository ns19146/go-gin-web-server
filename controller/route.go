package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func StartWebServer() {
	r := gin.Default()
	r.GET("/", Hello)

	r.Run()
}

func Hello(_ *gin.Context) {
	fmt.Println("Hello World")
}
