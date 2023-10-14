package handler

import (
	"encoding/json"
	"log"
	"os"

	"github.com/ClearSyntax618/fiber_exchange-rate/models"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

// Obtengo la API KEY
func apiKeyCall() string {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	API_KEY := os.Getenv("APP_ID")

	return API_KEY
}

// Se ejecuta cuando visito /historical
func GetHistorical(c *fiber.Ctx) error {
	/*
		D - Done | U - Undone
		To Do:
			1. Obtener la fecha de la moneda. (D)
			2. Configurar llamada a la API. (D)
			3. Hacer la llamada a la API con la fecha. (D)
			4. Retornar lo que obtengo de la API. (D)
	*/
	// Obtener la fecha de la moneda.
	var date models.Date
	c.BodyParser(&date)

	// Configurar llamada a la API
	url := "https://openexchangerates.org/api/historical/" + date.Date + ".json?app_id=" + apiKeyCall()

	// Hacer la peticion a la url
	agent := fiber.Get(url)
	statusCode, body, errs := agent.Bytes()

	if len(errs) > 0 {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"errs": errs,
		})
	}

	var res fiber.Map
	err := json.Unmarshal(body, &res)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"err": err,
		})
	}

	// Mostrar response de la API
	return c.Status(statusCode).JSON(res)
}

func GetLatest(c *fiber.Ctx) error {
	/*
		To Do:
			1. Configurar llamada a la API.
			2. Obtener response de la API.
			3. Enviar data al front.
	*/
	return c.SendString("Hola Mundo!")
}
