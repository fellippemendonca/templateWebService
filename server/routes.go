package server

import (
	"github.com/fellippemendonca/templateWebService/services"
	"github.com/gin-gonic/gin"

	users "github.com/fellippemendonca/templateWebService/domains/users/controllers"
)

// InitRoutes routes
func InitRoutes(router *gin.Engine, svc *services.Services) {
	v1 := router.Group("/v1")
	{
		v1.GET("users", users.GetAllUsers(svc))
		v1.GET("users/:id", users.GetUserByID(svc))
	}
}
