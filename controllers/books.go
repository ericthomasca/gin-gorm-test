package controllers

import (
	"example/gin-gorm-test/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// FindBooks gets all books. (GET /books)
func FindBooks(c *gin.Context) {
	var books []models.Book
	models.DB.Find(&books)

	c.JSON(http.StatusOK, gin.H{"data": books})
}

// FindBook finds a book by ID. (GET /books/:id)
func FindBook(c *gin.Context) {
	var book models.Book

	err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": book})
}

// CreateBook creates a new book. (POST /books)
func CreateBook(c *gin.Context) {
	// Validate input
	var input models.CreateBookInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create book
	book := models.Book{
		Title:  input.Title,
		Author: input.Author,
	}
	models.DB.Create(&book)

	c.JSON(http.StatusOK, gin.H{"data": book})
}

// UpdateBook update a book. (PATCH /books/:id)
func UpdateBook(c *gin.Context) {
	// Get models if exists
	var book models.Book
	
	err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
	}

	// Validate input
	var input models.UpdateBookInput

	err = c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&book).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": book})
}

// DeleteBook deletes a book. (DELETE /books/:id)
func DeleteBook(c *gin.Context) {
	// Get model if exist
	var book models.Book

	err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&book)

	c.JSON(http.StatusOK, gin.H{"data": true})
}

