package database

import (
	"first_go_api/config"
	"fmt"
	"strconv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	p := config.Get("DB_PORT")
	host := config.Get("DB_HOST")
	user := config.Get("DB_USER")
	password := config.Get("DB_PASSWORD")
	name := config.Get("DB_NAME")

	port, _ := strconv.ParseUint(p, 10, 32)
	dns := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, name)
	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database")
	}

	fmt.Println("Connection Opened to Database")
	return db, err
}
