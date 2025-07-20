package main

import (
	"fmt"
	"github.com/land1725/database-assignment/gorm-advanced/migrations"
	"github.com/land1725/database-assignment/gorm-advanced/models"
	"github.com/land1725/database-assignment/gorm-advanced/queries"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"math/rand"
)

var db *gorm.DB

func init() {
	username := "root"     //账号
	password := "land1725" //密码
	host := "127.0.0.1"    //数据库地址，可以是Ip或者域名
	port := 3306           //数据库端口
	Dbname := "gorm"       //数据库名
	timeout := "10s"       //连接超时，10秒

	// root:root@tcp(127.0.0.1:3306)/gorm?
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=%s", username, password, host, port, Dbname, timeout)
	//连接MYSQL, 获得DB类型实例，用于后面的数据库读写操作。
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		//SkipDefaultTransaction: true,
		//NamingStrategy: schema.NamingStrategy{
		//	//TablePrefix:   "t_",
		//	SingularTable: true,
		//	NoLowerCase:   false,
		//},
	})
	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}
	// 连接成功
	fmt.Println(db)
}

// createUsers 创建5个测试用户
func createUsers(db *gorm.DB) []models.User {
	names := []string{"张三", "李四", "王五", "赵六", "钱七"}
	var users []models.User

	for _, name := range names {
		user := models.User{
			Name:      name,
			PostCount: 0, // 随机1-5篇文章
		}
		if err := db.Create(&user).Error; err != nil {
			log.Printf("创建用户失败: %v", err)
			continue
		}
		users = append(users, user)
		fmt.Printf("创建用户: %s (ID: %d)\n", name, user.ID)
	}

	return users
}

// createPosts 为每个用户创建一些文章
func createPosts(db *gorm.DB, users []models.User) []models.Post {
	titles := []string{
		"Go语言入门指南",
		"GORM使用技巧",
		"数据库设计原则",
		"Web开发最佳实践",
		"微服务架构解析",
		"算法与数据结构",
		"Linux系统管理",
		"Docker容器技术",
		"Kubernetes入门",
		"云原生应用开发",
	}

	contents := []string{
		"这是一篇关于Go语言基础知识的文章...",
		"本文将介绍GORM的高级用法...",
		"良好的数据库设计是应用成功的关键...",
		"现代Web开发需要遵循的一些原则...",
		"微服务架构如何提高系统可扩展性...",
		"算法是计算机科学的核心...",
		"Linux系统管理员需要掌握的基本技能...",
		"Docker改变了应用的部署方式...",
		"Kubernetes已成为容器编排的事实标准...",
		"云原生应用开发的新趋势...",
	}

	var posts []models.Post

	for _, user := range users {
		// 每个用户创建1到3篇文章
		postCount := rand.Intn(3) + 1
		for i := 0; i < postCount; i++ {
			title := titles[rand.Intn(len(titles))]
			content := contents[rand.Intn(len(contents))]
			status := "open"
			if rand.Intn(2) == 0 {
				status = "closed"
			}

			post := models.Post{
				Title:         title,
				Content:       content,
				UserID:        user.ID,
				CommentStatus: status,
			}

			if err := db.Create(&post).Error; err != nil {
				log.Printf("创建文章失败: %v", err)
				continue
			}
			posts = append(posts, post)
			fmt.Printf("为用户 %s 创建文章: %s (ID: %d)\n", user.Name, title, post.ID)
		}
	}

	return posts
}

// createComments 为文章创建评论
func createComments(db *gorm.DB, posts []models.Post) {
	comments := []string{
		"好文章！",
		"非常有帮助",
		"我不同意这个观点",
		"感谢分享",
		"能不能更详细一点？",
		"已收藏",
		"期待后续更新",
		"实践过了，确实有效",
		"有更好的解决方案吗？",
		"这个技术已经过时了",
	}

	for _, post := range posts {
		// 每篇文章创建0到5条评论
		commentCount := rand.Intn(6)
		for i := 0; i < commentCount; i++ {
			content := comments[rand.Intn(len(comments))]

			comment := models.Comment{
				Content: content,
				PostID:  post.ID,
			}

			if err := db.Create(&comment).Error; err != nil {
				log.Printf("创建评论失败: %v", err)
				continue
			}
			fmt.Printf("为文章 %s 创建评论: %s (ID: %d)\n", post.Title, content, comment.ID)
		}
	}
}
func main() {
	defer func() {
		if db != nil {
			sqlDB, _ := db.DB()
			err := sqlDB.Close()
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("close database success")
			}
		}
	}()
	err := migrations.AutoMigration(db)
	if err != nil {
		fmt.Println(err.Error())
	}
	//hooks.RegisterPostHooks(db)
	//hooks.RegisterCommentHooks(db)
	// 清空现有数据
	//db.Unscoped().Where("1 = 1").Delete(&models.Comment{})
	//db.Unscoped().Where("1 = 1").Delete(&models.Post{})
	//db.Unscoped().Where("1 = 1").Delete(&models.User{})

	// 生成测试数据
	users := createUsers(db)
	posts := createPosts(db, users)
	createComments(db, posts)

	fmt.Println("测试数据生成完成！")

	user := queries.QueryUserPostComment(db, 4)
	fmt.Println(user)

	popularPost := queries.QueryPopularPost(db)
	fmt.Printf("Popular Post ID: %d\n", popularPost.ID)

	fmt.Printf("要删除所有评论的 文章 Post ID: %d\n", popularPost)
	//批量删除
	var delComment []models.Comment
	//delComment.PostID = popularPost.ID
	db.Debug().Where("post_id = ?", popularPost.ID).Find(&delComment)
	if err := db.Debug().Delete(&delComment).Error; err != nil {
		fmt.Printf("删除评论失败: %v\n", err)
	}
	//////单条删除
	//var delComment models.Comment
	//delComment.Content = "我不同意这个观点"
	//db.Debug().Where(delComment).First(&delComment)
	//
	//if err := db.Debug().Delete(&delComment).Error; err != nil {
	//	fmt.Printf("删除评论失败: %v\n", err)
	//}
}
