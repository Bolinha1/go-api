package domain

import (
	"database/sql"
	"go-api/config/db"

	_ "github.com/go-sql-driver/mysql"
)

func GetUser() {
	dsn := db.GetDsn()
	db, err := sql.Open("mysql", dsn)

}
