package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"log"
)

type Params struct {
	DB      *gorm.DB
	Current uint
	PerPage uint
	OrderBy string
}

type PaginatorMeta struct {
	Total   uint
	PerPage uint
	Current uint
}

type Paginator struct {
	Meta *PaginatorMeta
	Data interface{}
}

func Paginate(p Params, models interface{}) *Paginator {
	var per uint = 10
	if p.PerPage != 0 {
		per = p.PerPage
	}
	var current uint = 1
	if p.Current != 0 {
		current = p.Current
	}
	var cnt uint
	p.DB.Find(models).Count(&cnt)

	offset := per*(current-1) + 1
	p.DB.Limit(p.PerPage).Offset(offset).Order(p.OrderBy).Find(models)
	meta := PaginatorMeta{
		Total:   cnt,
		PerPage: per,
		Current: current,
	}
	return &Paginator{
		Meta: &meta,
		Data: models,
	}
}

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
	x := db.Preload("User")
	paginator := Paginate(Params{
		DB:      x,
		PerPage: 1,
		Current: 1,
	}, &posts)

	fmt.Println(paginator.Meta)
}
