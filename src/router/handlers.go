package router

import (
	"github.com/gin-gonic/gin"
	"github.com/users-manager/src/models"
	"log"
)

func getAllUsersHandler(c *gin.Context) {
	userName := c.Request.Header.Get(authHeader)

	if userName == "" {
		//TODO handle empty user
		return
	}

	user := new(models.User)
	dbManager.Where(models.User{Name: userName}).First(user)

	if user == nil {
		//TODO handle user is absent
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
