package middleware

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

// SQLDB middleware
func SQLDB(key string, sqldb *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(key, sqldb)
		c.Next()
	}
}
