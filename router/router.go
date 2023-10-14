package router

import (
	"github.com/ClearSyntax618/fiber_exchange-rate/handler"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// Rutas exchange
	coin := app.Group("/")
	coin.Get("/historical", handler.GetHistorical)
	coin.Get("/latest", handler.GetLatest)

}
