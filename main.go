package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	controller "gin-book-mementos/controller"
)

func main() {
	router := gin.Default()

	// 静的ファイルのパスを指定
	router.Static("/assets", "./assets")
	router.Static("/view", "./view")

	router.LoadHTMLGlob("templates/*")

	// ルーティング定義
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Book Mementos",
		})
	})
	// 全ての本の情報を返す
	router.GET("/fetchAllBooksFromList", controller.FetchAllBooksFromList)

	// 本をDBヘ登録
	router.POST("/addBookToList", controller.AddBookToList)

	// DBから本を削除
	router.POST("/deleteBookFromList", controller.DeleteBookFromList)

	// 本の読了状態を変更する
	router.POST("/changeBookState", controller.ChangeBookState)

	router.Run(":8080")
}
