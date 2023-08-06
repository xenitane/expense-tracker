package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/xenitane/expense-tracker/server/controllers"
)

func RouteConfig(fiberApp *fiber.App) {

	fiberApp.Use("/health-check", controllers.HealthCheck) // health-check
	controllers.IncomeController(fiberApp.Group("/api/v1/income"))
	controllers.ExpenseController(fiberApp.Group("/api/v1/expense"))
	fiberApp.Use("*", controllers.NotFound) //404s
}
