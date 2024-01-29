package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/facundocarballo/microservices-todo-list/api/routes"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

const (
	PORT = "3001"
)

func GetDSN() string {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	return dbUser + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName
}

func main() {

	// .env file
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Error loading the env file. %s\n", err.Error())
		os.Exit(1)
	}

	db, err := sql.Open("mysql", GetDSN())
	if err != nil {
		fmt.Printf("Error connecting to the database. %s\n", err.Error())
		os.Exit(1)
	}
	defer db.Close()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "POST, GET, PUT, DELETE",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	routes.SetUp(app, db)

	app.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).SendString("Endpoint not found.")
	})

	app.Listen(":" + PORT)
}
