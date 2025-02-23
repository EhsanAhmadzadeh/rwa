package models

import (
	"database/sql"
	"log"
	"wa-service/config"

	_ "github.com/mattn/go-sqlite3"
)

// DB is a package-level variable that holds the database connection.
var DB *sql.DB

// ConnectDB connects to the database and assigns the connection to DB.
func ConnectDB() error {
	var err error
	DB, err = sql.Open("sqlite3", config.AppConfig.DB_PATH)
	if err != nil {
		return err
	}
	if err = DB.Ping(); err != nil {
		return err
	}
	log.Println("✅ Database connected successfully")
	return nil
}
