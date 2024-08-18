package app

import (
	"database/sql"
	_ "github.com/golang-migrate/migrate/database/mysql"
	"log"
	"time"
)

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "joko:akuanakhebat@tcp(localhost:3306)/contact_management_go")
	if err != nil {
		log.Fatal(err.Error())
	}

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(time.Hour)
	db.SetConnMaxIdleTime(10 * time.Minute)
	return db
}
