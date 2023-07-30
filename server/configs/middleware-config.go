package configs

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/xenitane/expense-tracker/server/variables"
)

func MiddlewareConfig() {
	variables.App.Use(cors.New(cors.Config{
		ExposeHeaders: "Link",
		MaxAge:        300,
		AllowOrigins:  strings.Join([]string{"https://*", "https:/*"}, ","),
		AllowMethods: strings.Join([]string{
			fiber.MethodGet,
			fiber.MethodPost,
			fiber.MethodPut,
			fiber.MethodPatch,
			fiber.MethodDelete,
			fiber.MethodOptions,
			fiber.MethodHead,
		}, ","),
		AllowHeaders:     "*",
		AllowCredentials: false,
	})) // cors
	variables.App.Use(compress.New(compress.Config{Level: compress.LevelBestSpeed})) // compress
	variables.App.Use(helmet.New())                                                  // helmet
	variables.App.Use(logger.New())                                                  // logger
	variables.App.Use(func(c *fiber.Ctx) error {
		c.Accepts("application/json")
		c.AcceptsCharsets("utf-8")
		return c.Next()
	})
}
