package controller

import (
	"gin-book-mementos/models"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// 全ての本の情報を返す
func FetchAllBooksFromList(c *gin.Context) {
	db, err := gorm.Open("mysql", "user1:Password_01@tcp(docker.for.mac.localhost:3306)/books?charset=utf8&parseTime=True")
	if err != nil {
		log.Println(err)
	}

	books := []models.Book{}

	db.Order("book_id asc").Find(&books)

	c.JSON(200, books)

	defer db.Close()
}

// 指定された本の情報を返す
func FetchBookFromList(c *gin.Context) {

}

// データベースに登録された本を一つ見つける
func FindBookFirst(title string) models.Book {
	var book models.Book

	db, err := gorm.Open("mysql", "user1:Password_01@tcp(docker.for.mac.localhost:3306)/books?charset=utf8&parseTime=True")
	if err != nil {
		log.Println(err)
	}

	defer db.Close()
	if err := db.Where("title = ?", title).First(&book).Error; err != nil {
		log.Println(err)
	}
	return book
}

// 本がすでにDBに登録されているかどうか確認、tureなら登録済み、falseなら未登録
func JudgeBookInList(title string) bool {
	bookTitle := title

	if b := FindBookFirst(bookTitle); b.Title == bookTitle {
		return true
	} else {
		return false
	}
}

// 本をDBヘ登録
func AddBookToList(c *gin.Context) {
	db, err := gorm.Open("mysql", "user1:Password_01@tcp(docker.for.mac.localhost:3306)/books?charset=utf8&parseTime=True")
	if err != nil {
		log.Println(err)
	}

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

	// DBに本がすでに登録されていればstatus:409を返す、登録されていなければDBに本を登録する
	if JudgeBookInList(bookTitle) {
		c.Status(409)
	} else {
		db.Create(&book)
	}

	defer db.Close()
}

// DBから本を削除
func DeleteBookFromList(c *gin.Context) {
	db, err := gorm.Open("mysql", "user1:Password_01@tcp(docker.for.mac.localhost:3306)/books?charset=utf8&parseTime=True")
	if err != nil {
		log.Println(err)
	}

	bookID, err := strconv.Atoi(c.PostForm("bookID"))
	if err != nil {
		log.Println(err)
	}

	book := []models.Book{}

	db.Delete(&book, bookID)

	defer db.Close()
}

// 本の読了状態を変更する
func ChangeBookState(c *gin.Context) {

}
