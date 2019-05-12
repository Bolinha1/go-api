package domain

import (
	"database/sql"
	"go-api/config/db"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func GetUser() []User {
	dsn := db.GetDsn()
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		log.Print(err.Error())
	}
	defer db.Close()

	results, err := db.Query("SELECT id, name, email FROM user")

	if err != nil {
		panic(err.Error())
	}

	for results.Next() {

		var users User

		err = results.Scan(&users.ID, &users.Name, &users.Email)
		if err != nil {
			panic(err.Error())
		}

		UserModel = append(UserModel, User{ID: users.ID, Name: users.Name, Email: users.Email})
	}
	return UserModel

}
