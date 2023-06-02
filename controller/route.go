package controller

import (
	"github.com/gin-gonic/gin"
)

func StartWebServer() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*.html")
	r.POST("/insert", InsertTable)
	r.GET("/create", CreateTable)
	r.GET("/show", ShowTables)
	r.GET("/upload", UploadCsv)
	r.POST("/upload", OpenCsv)

	return r
}
