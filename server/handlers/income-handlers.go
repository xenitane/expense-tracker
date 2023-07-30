package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/xenitane/expense-tracker/server/models"
)

func AddIncomeHandler(c *fiber.Ctx) error {
	body := new(models.IncomeModel)
	err := c.BodyParser(&body)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"msg":     "can't parse json",
			"error":   err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "success",

		"data": fiber.Map{
			"income": body,
		},
	})
}
