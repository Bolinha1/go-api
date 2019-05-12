package db

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

//GetDsn
func GetDsn() string {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	database := os.Getenv("DB_NAME")

	return setDsn(user, pass, host, database)
}

func setDsn(user string, pass string, host string, db string) string {
	var dsn string
	dsn += user
	dsn += ":"
	dsn += pass
	dsn += "@tcp("
	dsn += host
	dsn += ")/"
	dsn += db
	return dsn
}
