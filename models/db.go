package models

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB

func InitDB(dotenvPath ...string) {
	var conn *gorm.DB
	var err error

	conn, err = gorm.Open("sqlite3", ":memory:")

	if err != nil {
		log.Panic(err.Error())
	}

	conn.AutoMigrate(&Users{}, &Locations{}, &LocationImages{})

	DB = conn
}

func CloseDB() {
	DB.Close()
}
