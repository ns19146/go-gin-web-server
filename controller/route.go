package controller

import (
	"github.com/gin-gonic/gin"
)

func StartWebServer() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*.html")
	r.POST("/insert", test)
	r.GET("/create", CreateTable)
	r.GET("/show", ShowTables)
	r.GET("/upload", UploadCsv)
	r.POST("/upload", OpenCsv)
	r.Run()
}
