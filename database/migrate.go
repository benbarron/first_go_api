package database

import "first_go_api/models"

func Migrate() {
	db, err := Connect()
	if err != nil {
		panic("Could not connect to database")
	}
	db.AutoMigrate(&models.User{})
}
