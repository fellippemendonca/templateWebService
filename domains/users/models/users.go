package users

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/fellippemendonca/templateWebService/services"
)

// User struct
type User struct {
	ID        int
	FirstName string
	LastName  string
}

// GetAllUsers model method
func GetAllUsers(svc *services.Services) []User {
	queryOne := "select ID, FIRST_NAME, LAST_NAME from profile"
	rows, err := svc.Mysql.Query(queryOne)

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	users := []User{}

	for rows.Next() {
		user := User{}
		err := rows.Scan(&user.ID, &user.FirstName, &user.LastName)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return users
}

// GetUserByID model method
func GetUserByID(svc *services.Services, id int) User {
	query := "SELECT ID, FIRST_NAME, LAST_NAME FROM profile WHERE ID = ? ;"
	var user User
	row := svc.Mysql.QueryRow(query, id)
	err := row.Scan(&user.ID, &user.FirstName, &user.LastName)
	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		panic(err)
	case nil:
		return user
	default:
		panic(err)
	}
}
