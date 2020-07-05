package users

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/fellippemendonca/templateWebService/internal/services/databases/mysql"
)

type node struct {
	id        int
	firstName string
	lastName  string
}

// GetUsers get all users
func GetUsers(c *gin.Context) {
	// user := c.Params.ByName("name")

	//users := []node{}
	users := []string{}

	var (
		id        int
		firstName string
		lastName  string
	)

	// queryOne := "select ID, FIRST_NAME, LAST_NAME from profile where ID = ?"
	// rows, err := db.Query(queryOne, 1)

	queryAll := "select ID, FIRST_NAME, LAST_NAME from profile"
	rows, err := mysql.DB.Query(queryAll)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&id, &firstName, &lastName)
		if err != nil {
			log.Fatal(err)
		}
		// n := node{id: id, firstName: firstName, lastName: lastName}
		// users = append(users, n)
		users = append(users, firstName)

	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	log.Println(users)
	c.JSON(200, gin.H{
		"users": users,
	})
}
