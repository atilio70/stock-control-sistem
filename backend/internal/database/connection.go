package database

import (
	"log"

	"github.com/atilio70/stock-control-sistem/backend/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB is the global database connection
var DB *gorm.DB

// Connect initializes the MySQL database connection
func Connect() {
	//Change usernames, password, host, and database name accordingly
	dsn := "host=localhost user=usuario_app password=miNuevaPassword dbname=gestion_db port=5432 sslmode=disable TimeZone=America/Argentina/Buenos_Aires"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}

	DB = db

	//Auto-migrate database tables
	err = db.AutoMigrate(
		&models.Product{},
		&models.Client{},
		&models.Supplier{},
		&models.Sale{},
	)
	if err != nil {
		log.Fatal("[error] failed to migrate database: ", err)
	}
}
