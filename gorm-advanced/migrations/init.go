package migrations

import (
	"github.com/land1725/database-assignment/gorm-advanced/models"
	"gorm.io/gorm"
)

func AutoMigration(db *gorm.DB) error {
	err := db.AutoMigrate(&models.User{}, &models.Post{}, &models.Comment{})
	if err != nil {
		return err
	} else {
		return nil
	}
}
