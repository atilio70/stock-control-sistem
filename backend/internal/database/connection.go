package database

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB is the global database connection
var DB *gorm.DB

// Connect initializes the MySQL database connection
func Connect() {
	//Change usernames, password, host, and database name accordingly
	dsn := "user:password@tcp(127.0.0.1.3306)/gestion_db?charset=utf8mb4&parseTime=Tru&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to the databse: ", err)
	}

	DB = db
}
