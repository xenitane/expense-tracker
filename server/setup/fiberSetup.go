package setup

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/xenitane/expense-tracker/server/config"
	"github.com/xenitane/expense-tracker/server/value"
)

func StartApplication() {
	value.FiberApp = fiber.New()

	config.MiddlewareConfig(value.FiberApp)
	config.RouteConfig(value.FiberApp)

	log.Printf("Starting the web server at port: %v", value.Env[value.ConstPort])

	if err := value.FiberApp.Listen(fmt.Sprintf(":%v", value.Env[value.ConstPort])); err != nil {
		log.Fatalf("There was an error: %v\n while starting the server at port: %v", err, value.Env[value.ConstPort])
	}
}
