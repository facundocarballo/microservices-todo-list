package routes

import (
	"database/sql"

	"github.com/facundocarballo/microservices-todo-list/api/handlers"
	infrastructure "github.com/facundocarballo/microservices-todo-list/infrastructure/task"
	"github.com/gofiber/fiber/v2"
)

func SetUpTaskRoutes(app *fiber.App, db *sql.DB) {
	taskRepository := infrastructure.NewMySqlTaskRepository(db)
	taskRoutes := app.Group("/task")

	taskRoutes.Post("/", func(c *fiber.Ctx) error {
		return handlers.CreateTask(c, db, taskRepository)
	})
	taskRoutes.Get("/", func(c *fiber.Ctx) error {
		return handlers.GetTasksFromUser(c, db, taskRepository)
	})
	taskRoutes.Post("complete", func(c *fiber.Ctx) error {
		return handlers.CompleteTask(c, db, taskRepository)
	})
	taskRoutes.Post("delete", func(c *fiber.Ctx) error {
		return handlers.DeleteTask(c, db, taskRepository)
	})
}
