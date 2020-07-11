package server

import (
	"fmt"
	"log"

	"github.com/fellippemendonca/templateWebService/server/middleware"
	"github.com/fellippemendonca/templateWebService/services/databases/mysql"
	"github.com/gin-gonic/gin"
)

// Run Server
func Run() {
	router := gin.New()

	fmt.Printf("Connecting MySQL Database... ")
	mysqlDb, err := mysql.Connect()
	if err != nil {
		log.Panic(err)
	}
	mysqlMid := middleware.SQLDB("MYSQL", mysqlDb)
	fmt.Println("OK")

	fmt.Printf("Setting middlewares... ")
	router.Use(mysqlMid)
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	fmt.Println("OK")

	fmt.Println("Setting routes... ")
	InitRoutes(router)
	fmt.Println("Setting routes... OK")

	router.Run(":3000")
}
