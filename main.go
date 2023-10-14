/*
EXERCISE:
----------------------------------------------------------
Create a service that gives the latest and historical exchange rate for the currency.
Please use three tier network architecture in this challenge.
You can also consider using pub / sub, worker mode to process the request to make it more scalable.
-----------------------------------------------------------
*/

package main

import (
	"github.com/ClearSyntax618/fiber_exchange-rate/router"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Inicializando servidor
	app := fiber.New()

	// Archivos estáticos
	app.Static("/", "./public/")
	app.Static("/latest", "./public/latest/")
	app.Static("/historical", "./public/historical/")

	// Rutas
	router.SetupRoutes(app)

	// Servidor en puerto:
	app.Listen(":3000")
}
