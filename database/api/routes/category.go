package routes

import (
	"database/sql"

	"github.com/facundocarballo/microservices-todo-list/api/handlers"
	infrastructure "github.com/facundocarballo/microservices-todo-list/infrastructure/category"
	"github.com/gofiber/fiber/v2"
)

func SetUpCategoryRoutes(app *fiber.App, db *sql.DB) {
	categoryRepository := infrastructure.NewMySqlCategoryRepository(db)
	categoryRoutes := app.Group("/category")

	categoryRoutes.Get("/", func(c *fiber.Ctx) error {
		return handlers.GetCategoriesFromUser(c, db, categoryRepository)
	})
	categoryRoutes.Post("/", func(c *fiber.Ctx) error {
		return handlers.CreateCategory(c, db, categoryRepository)
	})
	categoryRoutes.Delete("/:id", func(c *fiber.Ctx) error {
		return handlers.DeleteCategory(c, db, categoryRepository)
	})
}
