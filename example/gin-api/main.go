package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"log"
)

type User struct {
	gorm.Model
	Name  string
	Posts []Post `gorm:"foreignkey:UserID"`
}

type Post struct {
	gorm.Model
	Content string
	UserID  uint
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalln(err)
	}

	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=postgres password=secret sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()
	db.AutoMigrate(&User{}, &Post{})

	user := User{
		Name: "John Doe",
	}
	db.Create(&user)
	post := Post{
		Content: "Hello.",
		UserID:  user.ID,
	}
	db.Create(&post)

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	_ = r.Run() // listen and serve on 0.0.0.0:8080
}
