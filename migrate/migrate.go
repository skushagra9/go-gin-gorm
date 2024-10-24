package main

import (
	intializers "go-gin-gorm/initializers"
	"go-gin-gorm/models"
)

func init() {
	intializers.LoadEnvVariables()
	intializers.InitializeDB()
}

func main() {
	intializers.DB.AutoMigrate(&models.Book{})
}
