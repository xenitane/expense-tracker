package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/xenitane/expense-tracker/server/model"
	"github.com/xenitane/expense-tracker/server/repository"
	"github.com/xenitane/expense-tracker/server/value"
)

func AddExpense(c *fiber.Ctx) error {
	expense := new(model.ExpenseModel)
	if c.BodyParser(&expense) != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   "error while parsing data",
		})
	}
	if value.Validator.Struct(expense) != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   "data validation failed",
		})
	}
	insertedIncome, err := repository.SaveExpense(expense)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"err":     "an error occoured while saving the record",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "record saved successfully",
		"data": fiber.Map{
			"income": insertedIncome,
		},
	})
}

func GetExpenses(c *fiber.Ctx) error {
	expenses, err := repository.GetExpenses()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   "an error occoured while getting the records",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"expenses": expenses,
		},
	})
}
func GetExpense(c *fiber.Ctx) error {
	id := c.Params("id")
	expense, err := repository.GetExpenseByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"error":   "no such document exists(" + err.Error() + ")",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"expense": expense,
		},
	})
}
func DeleteExpense(c *fiber.Ctx) error {
	id := c.Params("id")
	err := repository.DeleteExpense(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   "an error occoured while deleting the record(" + err.Error() + ")c",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "record deleted successfully",
	})
}
