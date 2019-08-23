package main

import (
	"fmt"
	"github.com/dyoshikawa/paginate-gorm/paginator"
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
	User    User
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

	var posts []Post
	query := paginator.NewQuery(paginator.QueryParams{
		DB:      db.Preload("User"),
		Models:  &posts,
		Current: 2,
	})
	paginator := paginator.Paginate(query)

	fmt.Println(paginator.Meta)
}
