package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/xenitane/expense-tracker/server/model"
	"github.com/xenitane/expense-tracker/server/repository"
)

func AddIncomeHandler(c *fiber.Ctx) error {
	income := new(model.IncomeModel)
	err := c.BodyParser(&income)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "can't parse json, either the data passed is invalid or incomplete",
			"error":   err.Error(),
		})
	}
	if income.Title == nil || income.Category == nil || income.Date == nil || income.Description == nil {
		errorString := ""
		sep := ""
		if income.Title == nil {
			errorString += sep + "title"
			sep = ", "
		}
		if income.Category == nil {
			errorString += sep + "category"
			sep = ", "
		}
		if income.Date == nil {
			errorString += sep + "date"
			sep = ", "
		}
		if income.Description == nil {
			errorString += sep + "description"
		}
		errorString += " is/are missing"
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "all the fields are required",
			"error":   errorString,
		})
	}

	if *(income.Amount) <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "amount must be a positive number",
			"error":   "invalid value error",
		})
	}

	insertedIncome, err := repository.SaveIncome(income)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"err":     err.Error(),
			"message": "An error occoured while inserting the record",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "success",
		"data": fiber.Map{
			"income": insertedIncome,
		},
	})
}

func GetIncomes(c *fiber.Ctx) error {

	return nil
}
