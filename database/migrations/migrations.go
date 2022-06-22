package migrations

import (
	"github.com/andrewscarlos/golang-book-api/models"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(models.Book{})
}
