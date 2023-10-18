package main

import (
	"github.com/ClearSyntax618/fiber_exchange-rate/router"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Inicializando servidor
	app := fiber.New()

	// Rutas
	router.SetupRoutes(app)

	// Servidor en puerto:
	app.Listen(":3000")
}
