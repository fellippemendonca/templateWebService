package server

import (
	"github.com/fellippemendonca/templateWebService/internal/server/router"
	"github.com/fellippemendonca/templateWebService/internal/services/databases/mysql"
)

// Init Server
func Init() {
	mysql.Connect()
	router := router.Init()
	router.Run(":3000")
}
