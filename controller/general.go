package controller

import (
	"net/http"

	"github.com/matisszilard/gondol/model"
	store "github.com/matisszilard/gondol/store"

	gin "gopkg.in/gin-gonic/gin.v1"
)

// HelloWorld hey
func RootController(c *gin.Context) {
	c.String(http.StatusOK, "Hello World!")
}

// PostUser create a new user
func PostUser(c *gin.Context) {
	in := &model.User{}
	err := c.Bind(in)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	user := &model.User{}
	user.Email = in.Email

	var id string
	id, err = store.CreateUser(user)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	user.ID = id
	c.JSON(http.StatusOK, user)
}

// GetUser return the user from the database base on "id"
func GetUser(c *gin.Context) {
	user, err := store.GetUser(c.Param("id"))

	if err != nil {
		c.String(http.StatusNotFound, "Cannot find user. %s", err)
		return
	}

	c.JSON(http.StatusOK, user)
}

// GetUsers return all users from the database
func GetUsers(c *gin.Context) {
	users, err := store.GetUsers()

	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, users)
}
