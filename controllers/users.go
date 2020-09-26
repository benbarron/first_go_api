package controllers

import (
	"first_go_api/database"
	"first_go_api/models"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func CreateUser(c *fiber.Ctx) error {
	db, err := database.Connect()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": "Error",
		})
	}

	user := new(models.User)
	if c.BodyParser(user) != nil || user.HashPassword() != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": "Error",
		})
	}

	db.Create(&user)
	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"status":  http.StatusCreated,
		"message": "user created",
		"user":    user,
	})
}

func GetAllUsers(c *fiber.Ctx) error {
	db, err := database.Connect()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	var users []models.User
	db.Find(&users)
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status": http.StatusOK,
		"users":  users,
	})
}

func DeleteUser(c *fiber.Ctx) error {
	db, err := database.Connect()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	var user models.User

	db.First(&user, c.Params("id"))
	db.Delete(&user)

	return c.Status(http.StatusAccepted).JSON(fiber.Map{
		"status":  http.StatusAccepted,
		"message": "User deleted",
	})
}
