package users

import (
	"database/sql"
	"strconv"

	users_v2_models "github.com/fellippemendonca/templateWebService/domains/users/v2/models"
	"github.com/gin-gonic/gin"
)

// GetUsers get all users
func GetUsers(c *gin.Context) {
	mysql := c.Keys["MYSQL"].(*sql.DB)

	var userID int
	userID, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		userID = 0
	}

	var user users_v2_models.User
	user = users_v2_models.GetUserByID(mysql, userID)

	c.JSON(200, user)
}
