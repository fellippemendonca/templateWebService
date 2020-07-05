package router

import (
	users "github.com/fellippemendonca/templateWebService/internal/modules/users/controllers"
	"github.com/gin-gonic/gin"
)

// Init routes
func Init() *gin.Engine {
	// router := gin.Default()
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	/*
		authorized := r.Group("/")
		{
			authorized.POST("/login", loginEndpoint)
			authorized.POST("/submit", submitEndpoint)
			authorized.POST("/read", readEndpoint)
			// nested group
			testing := authorized.Group("testing")
			testing.GET("/analytics", analyticsEndpoint)
		}
	*/
	v1 := router.Group("/v1")
	{
		v1.GET("/user/:name", users.GetUsers)
	}
	return router
}
