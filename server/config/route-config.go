package config

import (
	"github.com/xenitane/expense-tracker/server/controllers"
	"github.com/xenitane/expense-tracker/server/values"
)

func RouteConfig() {

	values.App.Use("/health-check", controllers.HealthCheck) // health-check
	values.App.Use("api/v1/income", controllers.IncomeController())
	values.App.Use("*", controllers.NotFound) //404s
}
