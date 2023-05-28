package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/render-examples/go-gin-web-server/models"
	"net/http"
)

func dbInit() *gorm.DB {
	dsn := "host=dpg-chdin0bhp8u3v70u25og-a port=5432 user=nittc2023_j5exp_g2 dbname=nittc2023_j5exp_g2 password=uEeRkwJRnQufgEbbF3EFnrJUm0BDJRzP"
	db, err := gorm.Open("postgres", dsn)
	if err != nil {
		panic("failed to connect database")
	}

	return db
}

func test(c *gin.Context) {
	var model models.TestModel
	db := dbInit()
	if err := c.ShouldBindJSON(&model); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	db.Create(&model)
}
