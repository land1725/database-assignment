package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title         string
	Content       string
	UserID        uint
	Comments      []Comment `gorm:"foreignKey:PostID"`
	CommentStatus string    `gorm:"default:'有评论'"`
	CommentCount  int       `gorm:"default:0"`
}

func (post *Post) AfterCreate(tx *gorm.DB) error {
	return tx.Debug().Model(&User{}).
		Where("id = ?", post.UserID).
		UpdateColumn("post_count", gorm.Expr("post_count + 1")).Error
}
