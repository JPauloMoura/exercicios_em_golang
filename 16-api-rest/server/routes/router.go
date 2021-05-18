package routes

import (
	"github.com/JPaulo-Moura/16-api-rest/controllers"
	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1") //seta a rota padrão http://locahost/api/v1
	{
		books := main.Group("books")
		{
			books.GET("/", controllers.ShowBooks)
			books.GET("/:id", controllers.ShowBook)
			books.POST("/", controllers.CreateBook)
			books.PUT("/", controllers.UpdateBook)
			books.DELETE("/:id", controllers.DeleteBook)
		}
	}

	return router //GET http://locahost/api/v1/books
}
