package controller

import (
	"gin-book-mementos/models"
	"github.com/gin-gonic/gin"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// 全ての本の情報を返す
func FetchAllBooksFromList(c *gin.Context) {

}

// 指定された本の情報を返す
func FetchBookFromList(c *gin.Context) {

}

// 本をDBヘ登録
func AddBookToList(c *gin.Context) {
	db, _ := gorm.Open("mysql", "user1:Password_01@tcp(localhost:3306)/books?charset=utf8&parseTime=True")

	imageLink := c.PostForm("imageLink")
	bookTitle := c.PostForm("bookTitle")
	bookSubTitle := c.PostForm("bookSubTitle")
	authors := c.PostForm("authors")
	publisher := c.PostForm("publisher")
	publishedDate := c.PostForm("publishedDate")
	previewLink := c.PostForm("previewLink")

	var book = models.Book{
		ImageLink:     imageLink,
		Title:         bookTitle,
		SubTitle:      bookSubTitle,
		Authors:       authors,
		Publisher:     publisher,
		PublishedDate: publishedDate,
		PreviewLink:   previewLink,
		State:         false,
	}

	db.Create(&book)

	defer db.Close()
}

// DBから本を削除
func DeleteBookFromList(c *gin.Context) {

}

// 本の読了状態を変更する
func ChangeBookState(c *gin.Context) {

}
