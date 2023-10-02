package routes

import (
	"fmt"

	"github.com/RomainC75/postgres-test/controllers"
	"github.com/gin-gonic/gin"
)

func display(c *gin.Context) {
	fmt.Print("--> ROUTE MIDDELWARE \n")

	c.Next()
}

func middSetter(c *gin.Context) {
	c.Set("middValue", "myValue")
	c.Next()
}

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(display)

	grp1 :=
		r.Group("/book")
	{
		grp1.POST("/", controllers.AddBook)
		grp1.GET("/", middSetter, controllers.ListBooks)
		grp1.GET("/:id", controllers.GetBook)
		grp1.PUT("/:id", controllers.UpdateBook)
		grp1.DELETE("/:id", controllers.DeleteBook)
	}

	grp2 :=
		r.Group("/user")
	{
		grp2.POST("/signup", controllers.SignupUser)
		grp2.POST("/login", controllers.LoginUser)
		// grp2.GET("/", controllers.ListBooks)
		// grp2.GET("/:id", controllers.GetBook)
		// grp2.PUT("/:id", controllers.UpdateBook)
		// grp2.DELETE("/:id", controllers.DeleteBook)
	}
	return r
}
