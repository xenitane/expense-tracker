package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/xenitane/expense-tracker/server/service"
)

func ExpenseController(fiberApp *fiber.App) {
	fiberApp.Group("/api/v1/expense").Get("/", service.GetExpenses).Post("/", service.AddExpense).Get("/:id", service.GetExpense).Delete("/:id", service.DeleteExpense)

}
