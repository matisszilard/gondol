package router

import (
	"github.com/matisszilard/gondol/controller"
	gin "gopkg.in/gin-gonic/gin.v1"
)

// Serve the http server
func Serve() {
	router := gin.Default()

	router.GET("/", controller.RootController)

	users := router.Group("/users")
	{
		users.GET("", controller.GetUsers)
		users.POST("", controller.PostUser)

		users.GET("/:id", controller.GetUser)
	}

	router.Run(":8080")
}
