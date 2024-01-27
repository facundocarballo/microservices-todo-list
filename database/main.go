package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/facundocarballo/microservices-todo-list/domain"
	infrastructure "github.com/facundocarballo/microservices-todo-list/infrastructure/user"
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

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("¡Hola, Fiber!")
	})

	app.Post("/user", func(c *fiber.Ctx) error {
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
	})

	app.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).SendString("Endpoint not found.")
	})

	app.Listen(":" + PORT)
}
