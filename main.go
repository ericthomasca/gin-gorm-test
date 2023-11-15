package main

import (
	"example/gin-gorm-test/controllers"
	"example/gin-gorm-test/models"
	
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.SetTrustedProxies([]string{"localhost"})

	models.ConnectDatabase()

	r.GET("/books", controllers.FindBooks)
	r.GET("/books/:id", controllers.FindBook)
	r.POST("/books", controllers.CreateBook)
	r.PATCH("/books/:id", controllers.UpdateBook)
	r.DELETE("books/:id", controllers.DeleteBook)

	r.Run(":8099")
}
