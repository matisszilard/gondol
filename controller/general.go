package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HelloWorld hey
func HelloWorld(c *gin.Context) {
	c.String(http.StatusOK, "Hello World!")
}
