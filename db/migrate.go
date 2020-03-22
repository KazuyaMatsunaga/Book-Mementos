package main

import (
	"gin-book-mementos/models"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func main() {
	db, _ := gorm.Open("mysql", "user1:Password_01@tcp(localhost:3306)/books?charset=utf8&parseTime=True")
	db.CreateTable(&models.Book{})
}
