package models

import (
	"fmt"
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	Content string
	PostID  uint `gorm:"index"`
}

func (comment *Comment) AfterCreate(tx *gorm.DB) error {
	return tx.Debug().Model(&Post{}).
		Where("id = ?", comment.PostID).
		UpdateColumn("comment_count", gorm.Expr("comment_count + 1")).Error
}

func (c *Comment) AfterDelete(tx *gorm.DB) error {
	fmt.Println("进入删除，执行了一次")
	return tx.Debug().Model(&Post{}).
		Where("id = ?", c.PostID).
		Updates(map[string]interface{}{
			"comment_count": gorm.Expr("comment_count - ?", 1), // 关键改动：使用实际删除行数
			"comment_status": gorm.Expr(`
                CASE WHEN (comment_count - ?) <= 0 THEN '无评论'
                ELSE comment_status END
            `, tx.Statement.RowsAffected),
		}).Error
}
