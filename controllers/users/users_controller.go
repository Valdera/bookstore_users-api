package users

import (
	"github.com/Valdera/bookstore_users-api/domain/users"
	"github.com/Valdera/bookstore_users-api/services"
	"github.com/Valdera/bookstore_users-api/utils/errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUser(c *gin.Context) {

	c.String(http.StatusNotImplemented, "implement me!")

}

func CreateUser(c *gin.Context) {
	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, err := services.CreateUser(user)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusCreated, result)
}

func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!")
}
