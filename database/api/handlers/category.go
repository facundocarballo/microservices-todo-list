package handlers

import (
	"database/sql"
	"strconv"

	"github.com/facundocarballo/microservices-todo-list/api/utils"
	"github.com/facundocarballo/microservices-todo-list/domain"
	domain_task "github.com/facundocarballo/microservices-todo-list/domain/task"
	infrastructure "github.com/facundocarballo/microservices-todo-list/infrastructure/category"
	"github.com/gofiber/fiber/v2"
)

func GetCategoriesFromUser(
	c *fiber.Ctx,
	db *sql.DB,
	categoryRepository infrastructure.CategoryRepository,
) error {
	userIdString := c.Query("user_id", "")
	if userIdString == "" {
		return c.Status(fiber.StatusBadRequest).SendString("'user_id' parameter not found.")
	}

	userId, err := strconv.Atoi(userIdString)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("'user_id' not a number.")
	}

	categories, err := categoryRepository.Get(&domain.User{Id: userId})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Cannot get all the categories for this user.")
	}

	c.Set(utils.CONTENT_TYPE, utils.APP_JSON)
	return c.Status(fiber.StatusOK).JSON(categories)
}

func CreateCategory(
	c *fiber.Ctx,
	db *sql.DB,
	categoryRepository infrastructure.CategoryRepository,
) error {
	category := domain_task.BodyToCategory(c.Body())
	if category == nil {
		return c.Status(fiber.StatusBadRequest).SendString(utils.CATEGORY_BAD_RQUEST_MSG)

	}

	category, err := categoryRepository.Create(category)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Cannot create the category, check the body request.")
	}

	c.Set(utils.CONTENT_TYPE, utils.APP_JSON)
	return c.Status(fiber.StatusOK).JSON(category)
}

func DeleteCategory(
	c *fiber.Ctx,
	db *sql.DB,
	categoryRepository infrastructure.CategoryRepository,
) error {
	return nil
}
