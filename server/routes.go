package server

import (
	users_v1 "github.com/fellippemendonca/templateWebService/domains/users/v1/controllers"
	users_v2 "github.com/fellippemendonca/templateWebService/domains/users/v2/controllers"
	"github.com/gin-gonic/gin"
)

// InitRoutes routes
func InitRoutes(router *gin.Engine) {

	v1 := router.Group("/v1")
	{
		v1.GET("users", users_v1.GetUsers)
	}

	v2 := router.Group("/v2")
	{
		v2.GET("users/:id", users_v2.GetUsers)
	}
}
