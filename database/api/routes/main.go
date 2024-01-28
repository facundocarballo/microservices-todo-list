package routes

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
)

func SetUp(app *fiber.App, db *sql.DB) {
	SetUpUserRoutes(app, db)
	SetUpTaskRoutes(app, db)
}
