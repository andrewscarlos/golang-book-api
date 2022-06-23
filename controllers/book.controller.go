package controllers

import (
	"github.com/andrewscarlos/golang-book-api/database"
	"github.com/andrewscarlos/golang-book-api/models"
	"github.com/gin-gonic/gin"
	"strconv"
)

func ShowBook(c *gin.Context) {
	id := c.Param("id")

	newId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Id has to be integer",
		})
		return
	}
	db := database.GetDatabase()
	var p models.Book
	err = db.First(&p, newId).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot find book " + err.Error(),
		})
		return
	}
	c.JSON(200, p)
}

func CreateBook(c *gin.Context) {
	db := database.GetDatabase()

	var p models.Book

	err := c.ShouldBindJSON(&p)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	err = db.Create(&p).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot create book: " + err.Error(),
		})
		return
	}

	c.JSON(200, p)
}

func ShowBooks(c *gin.Context) {
	db := database.GetDatabase()
	var books []models.Book
	err := db.Find(&books).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot found list of books: " + err.Error(),
		})
		return
	}
	c.JSON(200, books)
}
