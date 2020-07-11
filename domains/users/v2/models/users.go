package users

import (
	"database/sql"
	"fmt"
	"log"
)

// User struct
type User struct {
	ID        int
	FirstName string
	LastName  string
}

// GetUserByID model method
func GetUserByID(mysql *sql.DB, id int) User {
	queryOne := "select ID, FIRST_NAME, LAST_NAME from profile where ID = ?"

	fmt.Printf("0.0 -----------------> %s, %d\n", queryOne, id)
	rows, err := mysql.Query(queryOne, id)

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
	log.Println(users[0])
	return users[0]
}
