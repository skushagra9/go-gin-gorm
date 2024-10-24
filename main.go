package main

import (
	"go-gin-gorm/controllers"
	intializers "go-gin-gorm/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	intializers.LoadEnvVariables()
	intializers.InitializeDB()
}

func main() {
	r := gin.Default()
	r.GET("/books", controllers.GetBooks)
	r.POST("/books", controllers.PostBook)
	r.GET("/book/:id", controllers.GetBook)
	r.Run()
}
