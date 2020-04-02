package controller

import (
	"gin-book-mementos/models"
	"github.com/gin-gonic/gin"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// 全ての本の情報を返す
func FetchAllBooksFromList(c *gin.Context) {
	db, _ := gorm.Open("mysql", "user1:Password_01@tcp(docker.for.mac.localhost:3306)/books?charset=utf8&parseTime=True")

	books := []models.Book{}

	db.Order("book_id asc").Find(&books)

	c.JSON(200, books)

	defer db.Close()
}

// 指定された本の情報を返す
func FetchBookFromList(c *gin.Context) {

}

// 本がすでに登録されているかどうか確認
func JudgeBookInList(c *gin.Context) {

}

// 本をDBヘ登録
func AddBookToList(c *gin.Context) {
	db, _ := gorm.Open("mysql", "user1:Password_01@tcp(docker.for.mac.localhost:3306)/books?charset=utf8&parseTime=True")

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
	db, _ := gorm.Open("mysql", "user1:Password_01@tcp(docker.for.mac.localhost:3306)/books?charset=utf8&parseTime=True")

	bookID, _ := strconv.Atoi(c.PostForm("bookID"))

	book := []models.Book{}

	db.Delete(&book, bookID)

	defer db.Close()
}

// 本の読了状態を変更する
func ChangeBookState(c *gin.Context) {

}
