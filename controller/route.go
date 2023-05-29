package controller

import (
	"github.com/gin-gonic/gin"
)

func StartWebServer() {
	r := gin.Default()
	r.POST("/insert", test)
	r.GET("/create", CreateTable)
	r.GET("/show", ShowTables)
	r.Run()
}
