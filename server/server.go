package server

import (
	"log"

	"github.com/fellippemendonca/templateWebService/services"
	"github.com/fellippemendonca/templateWebService/services/databases/mysql"
	"github.com/gin-gonic/gin"
)

// Run Server
func Run() {

	// Connect to required services
	mysqlDb, err := mysql.Connect()
	if err != nil {
		log.Panic(err)
	}

	// Assign services connection to an injectable struct
	svc := services.Services{}
	svc.Mysql = mysqlDb

	// Create Router and middleware
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// Init routes and inject started services access
	InitRoutes(router, &svc)

	// Start service on assigned port
	router.Run(":3000")
}
