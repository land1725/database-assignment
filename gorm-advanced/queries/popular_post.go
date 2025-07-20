package queries

import (
	"github.com/land1725/database-assignment/gorm-advanced/models"
	"gorm.io/gorm"
)

func QueryPopularPost(db *gorm.DB) models.Post {
	var post models.Post
	err := db.Debug().Order("comment_count desc").First(&post).Error
	if err == nil {
		return post
	} else {
		return models.Post{}
	}
}
