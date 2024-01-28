package routes

import (
	"database/sql"

	"github.com/facundocarballo/microservices-todo-list/api/handlers"
	infrastructure "github.com/facundocarballo/microservices-todo-list/infrastructure/user"
	"github.com/gofiber/fiber/v2"
)

func SetUpUserRoutes(app *fiber.App, db *sql.DB) {
	userRepository := infrastructure.NewMySqlUserRepository(db)
	userRoutes := app.Group("/user")

	userRoutes.Post("/", func(c *fiber.Ctx) error {
		return handlers.CreateUser(c, db, userRepository)
	})
	userRoutes.Get("/:id", func(c *fiber.Ctx) error {
		return handlers.GetUserById(c, db, userRepository)
	})
}
