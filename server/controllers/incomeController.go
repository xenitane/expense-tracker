package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/xenitane/expense-tracker/server/service"
)

func IncomeController(fiberApp *fiber.App) {
	fiberApp.Group("/api/v1/income").Get("/", service.GetIncomes).Post("/", service.AddIncome).Get("/:id", service.GetIncome).Delete("/:id", service.DeleteIncome)

}
