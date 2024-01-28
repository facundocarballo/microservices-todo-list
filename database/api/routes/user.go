package routes

import (
	"database/sql"

	"github.com/facundocarballo/microservices-todo-list/api/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetUpUserRoutes(app *fiber.App, db *sql.DB) {
	userRoutes := app.Group("/user")

	userRoutes.Post("/", func(c *fiber.Ctx) error {
		return handlers.CreateUser(c, db)
	})
}
