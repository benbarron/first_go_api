package main

import (
	"first_go_api/config"
	"first_go_api/controllers"
	"first_go_api/database"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	api := app.Group("/api")
	database.Migrate()
	api.Get("/users", controllers.GetAllUsers)
	api.Post("/users", controllers.CreateUser)
	api.Delete("/users/:id", controllers.DeleteUser)
	log.Fatal(app.Listen(":" + config.Get("SERVER_PORT")))
}
