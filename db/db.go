package db

import (
	"fmt"
	"gin-api/config"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"go.mau.fi/whatsmeow/store/sqlstore"
)

var DB *sqlstore.Container

// InitDB initializes the database connection and update the global DB
func InitDB() {
	sqlContainer, err := sqlstore.New("sqlite3", fmt.Sprintf("file:%s?_foreign_keys=on", config.DATABASE_URL), nil)
	if err != nil {
		log.Println("Couldn't initialize the database: ", err)
	} else {

		DB = sqlContainer
	}

}
