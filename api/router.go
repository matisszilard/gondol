package api

import (
	"net/http"

	gin "gopkg.in/gin-gonic/gin.v1"
)

func index(c *gin.Context) {
	c.String(http.StatusOK, "Hello World!")
}

// Serve the http server
func Serve() {
	router := gin.Default()

	router.GET("/", index)

	router.Run(":8080")
}
