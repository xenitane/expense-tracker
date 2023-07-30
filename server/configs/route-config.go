package configs

import (
	"github.com/xenitane/expense-tracker/server/handlers"
	"github.com/xenitane/expense-tracker/server/variables"
)

func RouteConfig() {

	variables.App.Get("/health-check", handlers.HealthCheckHandler) // health-check

	v1api := variables.App.Group("/api/v1") // api/v1 router

	v1api.Post("/add-income", handlers.AddIncomeHandler)

	variables.App.Use("*", handlers.NotFoundHandler) //404s
}
