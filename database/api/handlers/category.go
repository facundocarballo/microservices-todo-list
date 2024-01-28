package handlers

import (
	"database/sql"

	infrastructure "github.com/facundocarballo/microservices-todo-list/infrastructure/category"
	"github.com/gofiber/fiber/v2"
)

// TODO: Implementear metodos
func GetCategoriesFromUser(
	c *fiber.Ctx,
	db *sql.DB,
	categoryRepository infrastructure.CategoryRepository,
) error {
	return nil
}

func CreateCategory(
	c *fiber.Ctx,
	db *sql.DB,
	categoryRepository infrastructure.CategoryRepository,
) error {
	return nil
}

func DeleteCategory(
	c *fiber.Ctx,
	db *sql.DB,
	categoryRepository infrastructure.CategoryRepository,
) error {
	return nil
}
