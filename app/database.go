package app

import (
	"database/sql"
	"fmt"
	_ "github.com/golang-migrate/migrate/database/mysql"
	"log"
	"time"
)

func NewDB(config *Config) *sql.DB {

	DbUser := config.env.GetString("DB_USER")
	DbPassword := config.env.GetString("DB_PASSWORD")
	DbHost := config.env.GetString("DB_HOST")
	DbPort := config.env.GetString("DB_PORT")
	DbDatabase := config.env.GetString("DB_DATABASE")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", DbUser, DbPassword, DbHost, DbPort, DbDatabase)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err.Error())
	}

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(time.Hour)
	db.SetConnMaxIdleTime(10 * time.Minute)
	return db
}
