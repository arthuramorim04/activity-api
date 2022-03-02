package migrations

import (
	"github.com/arthuramorim04/go-activity-api.git/models"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(models.Activity{})
	db.AutoMigrate(models.User{})
}
