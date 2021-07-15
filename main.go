package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv" // new
	"github.com/nickel-dime/fiber-todo/config"
	"github.com/nickel-dime/fiber-todo/routes"
)

func main() {
	app := fiber.New()

	// dotenv
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config.ConnectDB()

	// give response when at /
	setupRoutes(app)

	// Listen on server 8000 and catch error if any
	err = app.Listen(":27017")

	// handle error
	if err != nil {
		panic(err)
	}
}

func setupRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"message": "You are at the endpoint 😉",
		})
	})

	api := app.Group("/api")

	api.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"message": "You are at the api endpoint 😉",
		})
	})

	routes.TodoRoute(api.Group("/todos"))

}
