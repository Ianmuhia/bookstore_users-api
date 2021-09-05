package users

import (
	"bookstore_users-api/domain/users"
	"bookstore_users-api/services"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func CreateUser(c *gin.Context) {
	var user users.User
	bytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		//TODO:Handle Error
		return
	}
	err = json.Unmarshal(bytes, &user)
	if err != nil {
		//TODO:Handle json Error
		return
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		//TODO: return bad request to the caller.
		return
	}
	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		//TODO:Handle UserCreation Error
		return
	}``

	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "impliment me")

}

func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "impliment me")

}
