package queries

import (
	"fmt"
	"github.com/land1725/database-assignment/gorm-advanced/models"
	"gorm.io/gorm"
)

func QueryUserPostComment(db *gorm.DB, userID uint) models.User {
	var user models.User
	err := db.Debug().Preload("Posts").Preload("Posts.Comments").Where("id = ?", userID).First(&user).Error
	if err != nil {
		fmt.Println(err)
		return models.User{}
	} else {
		return user
	}
}
