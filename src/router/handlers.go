package router

import (
	"github.com/gin-gonic/gin"
	"github.com/users-manager/src/models"
	"log"
	"net/http"
)

func getAllUsersHandler(c *gin.Context) {
	userName := c.Request.Header.Get(authHeader)

	if userName == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized ",
		})
		return
	}

	user := new(models.User)
	dbManager.Where(models.User{Name: userName}).First(user)

	if user == nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized ",
		})
		return
	}

	var users []models.User
	dbManager.Where(&models.User{Role: "user"}).Find(&users)
	log.Printf("Users %+v", users)

	c.JSON(200, gin.H{
		"message": "getAll",
	})
}

func addUserHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "addUser",
	})
}

func removeUserHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "removeUser",
	})
}
