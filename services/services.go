package services

import "database/sql"

// Services injector
type Services struct {
	Mysql *sql.DB
}
