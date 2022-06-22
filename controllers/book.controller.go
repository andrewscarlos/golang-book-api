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
	var book models.Book
	err = db.First(&book, newId).Error()
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot find book " + err.Error(),
		})
		return
	}
	c.JSON(200, book)
}
