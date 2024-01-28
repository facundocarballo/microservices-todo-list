package handlers

import (
	"database/sql"
	"strconv"

	"github.com/facundocarballo/microservices-todo-list/api/utils"
	"github.com/facundocarballo/microservices-todo-list/domain"
	infrastructure "github.com/facundocarballo/microservices-todo-list/infrastructure/user"
	"github.com/gofiber/fiber/v2"
)

func CreateUser(
	c *fiber.Ctx,
	db *sql.DB,
	userRepository infrastructure.UserRepository,
) error {
	user := domain.BodyToUser(c.Body())
	if user == nil {
		return c.Status(fiber.StatusBadRequest).SendString("Body of the request have to contains the specif interface of a user.")
	}
	user, err := userRepository.Create(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Cant create the user properly, check the body request.")
	}
	c.Set(utils.CONTENT_TYPE, utils.APP_JSON)
	return c.Status(fiber.StatusAccepted).JSON(user)
}

func GetUserById(
	c *fiber.Ctx,
	db *sql.DB,
	userRepository infrastructure.UserRepository,
) error {
	userIdString := c.Params("id")
	if userIdString == "" {
		return c.Status(fiber.StatusBadRequest).SendString("'id' parameter not found.")
	}

	userId, err := strconv.Atoi(userIdString)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("'id' parameter not a number.")
	}

	user, err := userRepository.Find(userId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("User not found.")
	}

	c.Set(utils.CONTENT_TYPE, utils.APP_JSON)
	return c.Status(fiber.StatusOK).JSON(user)
}
