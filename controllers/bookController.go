package controllers

import (
	intializers "go-gin-gorm/initializers"
	"go-gin-gorm/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetBooks(c *gin.Context) {
	var books []models.Book
	intializers.DB.Find(&books)
	c.JSON(http.StatusOK, gin.H{"data": books})
}

// Find(&books) tells the ORM to retrieve all records from the books table in the database and store them in the books slice.
// The &books means you are passing a reference to the slice so that the function can populate it with the result.

func PostBook(c *gin.Context) {
	var input models.Book
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book := models.Book{Title: input.Title, Author: input.Author}
	intializers.DB.Create(&book)

	c.JSON(http.StatusOK, gin.H{"data": book})
}

func GetBook(c *gin.Context) {
	var book models.Book
	if err := intializers.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": book})
}
