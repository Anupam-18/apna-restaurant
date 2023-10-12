package database

import (
	"apna-restaurant/models"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func DBConfig() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error in loading env vars", err)
		return
	}

	dbHost := os.Getenv("db_host")
	userName := os.Getenv("db_user")
	dbName := os.Getenv("db_name")
	password := os.Getenv("db_pass")
	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, userName, dbName, password)
	db, err = gorm.Open(postgres.Open(dbURI), &gorm.Config{})
	if err != nil {
		fmt.Println("Error while establishing connection", err)
		return
	}
	db.AutoMigrate(
		&models.User{},
		&models.Menu{},
		&models.MenuItem{},
	)
}
func GetDB() *gorm.DB {
	return db
}
