package api

import (
	"github.com/matisszilard/gondol/controller"
	gin "gopkg.in/gin-gonic/gin.v1"
)

// Serve the http server
func Serve() {
	router := gin.Default()

	router.GET("/", controller.HelloWorld)

	router.Run(":8080")
}
