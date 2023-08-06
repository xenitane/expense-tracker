package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/xenitane/expense-tracker/server/service"
)

func IncomeController(incomeRouter fiber.Router) {
	incomeRouter.Get("/", service.GetIncomes).Post("/", service.AddIncome).Get("/:id", service.GetIncome).Delete("/:id", service.DeleteIncome)
}
