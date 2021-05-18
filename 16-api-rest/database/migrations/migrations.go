package migrations

import (
	"github.com/JPaulo-Moura/16-api-rest/models"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(models.Book{})
}
