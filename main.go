package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	app := fiber.New()

	// Dotenv config
	_err := godotenv.Load()

	if _err != nil {
		log.Fatal("Error loading .env file")
	}

	API_KEY := os.Getenv("APP_ID")
	url := "https://openexchangerates.org/api/latest.json?app_id=" + API_KEY

	// Get currencies from API
	app.Get("/", func(c *fiber.Ctx) error {
		// Starting the request
		agent := fiber.Get(url)
		statusCode, body, errs := agent.Bytes()

		if len(errs) > 0 {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"errs": errs,
			})
		}

		// Unmarshalling data
		var coin fiber.Map
		err := json.Unmarshal(body, &coin)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"err": err,
			})
		}

		// Sending JSON data to "/"
		return c.Status(statusCode).JSON(coin)
	})

	// Convert currencies (?)

	app.Listen(":3000")
}
