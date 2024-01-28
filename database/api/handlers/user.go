package handlers

import (
	"database/sql"

	"github.com/facundocarballo/microservices-todo-list/domain"
	infrastructure "github.com/facundocarballo/microservices-todo-list/infrastructure/user"
	"github.com/gofiber/fiber/v2"
)

func CreateUser(c *fiber.Ctx, db *sql.DB) error {
	user := domain.BodyToUser(c.Body())
	if user == nil {
		return c.Status(fiber.StatusBadRequest).SendString("Body of the request have to contains the specif interface of a user.")
	}
	userRepository := infrastructure.NewMySqlUserRepository(db)
	user, err := userRepository.Create(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Cant create the user properly, check the body request.")
	}
	c.Set("Content-Type", "application/json")
	return c.Status(fiber.StatusAccepted).JSON(user)
}
