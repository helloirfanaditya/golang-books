package main

import (
	Books "micros/controllers"
	"micros/database"

	"github.com/gin-gonic/gin"
)

func main() {
	database.SetDB()
	defer database.Close()
	r := gin.Default()
	r.GET("/books", Books.GetAllBooks)
	r.GET("/book", Books.FindBook)
	r.POST("/book/create", Books.CreateBook)
	r.POST("/book/update", Books.UpdateBook)
	r.POST("/book/delete", Books.DeleteBook)
	r.Run(":1234")
}
