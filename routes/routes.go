package routes

import (
	"github.com/RomainC75/postgres-test/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	grp1 :=
		r.Group("/book")
	{
		grp1.POST("/", controllers.AddBook)
		grp1.GET("/", controllers.ListBooks)
		grp1.GET("/:id", controllers.GetBook)
		grp1.PUT("/:id", controllers.UpdateBook)
		grp1.DELETE("/:id", controllers.DeleteBook)
	}
	return r
}
