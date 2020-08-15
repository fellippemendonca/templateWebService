package users

import (
	"strconv"

	models "github.com/fellippemendonca/templateWebService/domains/users/models"
	"github.com/fellippemendonca/templateWebService/services"
	"github.com/gin-gonic/gin"
)

// GetAllUsers get all users
func GetAllUsers(svc *services.Services) func(c *gin.Context) {
	return func(c *gin.Context) {
		var users []*models.User
		users = models.GetAllUsers(svc)
		c.JSON(200, users)
	}
}

// GetUserByID get user by ID
func GetUserByID(svc *services.Services) func(c *gin.Context) {
	return func(c *gin.Context) {
		var userID int
		userID, err := strconv.Atoi(c.Params.ByName("id"))
		if err != nil {
			userID = 0
		}

		var user *models.User
		user = models.GetUserByID(svc, userID)

		c.JSON(200, user)
	}
}
