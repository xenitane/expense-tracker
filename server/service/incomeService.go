package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/xenitane/expense-tracker/server/model"
	"github.com/xenitane/expense-tracker/server/repository"
	"github.com/xenitane/expense-tracker/server/value"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddIncome(c *fiber.Ctx) error {
	income := new(model.IncomeModel)
	if c.BodyParser(&income) != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   "error while parsing data",
		})
	}
	income.TrimSpace()
	if value.Validator.Struct(income) != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   "data validation failed",
		})
	}
	insertedIncome, err := repository.SaveIncome(income)
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

func GetIncomes(c *fiber.Ctx) error {
	incomes, err := repository.GetIncomes()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   "an error occoured while getting the records" + err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"incomes": incomes,
		},
	})
}
func GetIncome(c *fiber.Ctx) error {
	id := c.Params("id")
	income, err := repository.GetIncomeByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"error":   "no such document exists(" + err.Error() + ")",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"income": income,
		},
	})
}
func DeleteIncome(c *fiber.Ctx) error {
	id, err := primitive.ObjectIDFromHex(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   "an error occoured while deleting the record(" + err.Error() + ")",
		})
	}
	err = repository.DeleteIncome(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   "an error occoured while deleting the record(" + err.Error() + ")",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "record deleted successfully",
	})
}
