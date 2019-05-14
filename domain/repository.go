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

	userList := []User{}

	for results.Next() {

		var users User

		err = results.Scan(&users.ID, &users.Name, &users.Email)
		if err != nil {
			panic(err.Error())
		}

		userList = append(userList, User{ID: users.ID, Name: users.Name, Email: users.Email})
	}
	return userList

}

func GetUserID(id string) User {
	dsn := db.GetDsn()
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		log.Print(err.Error())
	}
	defer db.Close()

	results, err := db.Prepare("SELECT id, name, email FROM user WHERE id = (?)")

	if err != nil {
		panic(err.Error())
	}

	var users User

	err = results.QueryRow(id).Scan(&users.ID, &users.Name, &users.Email)
	if err != nil {
		panic(err.Error())
	}

	results.Close()

	userModel = User{ID: users.ID, Name: users.Name, Email: users.Email}
	return userModel
}

func CreateUser(name string, email string) {

	dsn := db.GetDsn()
	db, err := sql.Open("mysql", dsn)

	insert, err := db.Query("INSERT INTO user (name, email) VALUES ( '" + name + "', '" + email + "' )")

	if err != nil {
		panic(err.Error())
	}

	defer insert.Close()
}
