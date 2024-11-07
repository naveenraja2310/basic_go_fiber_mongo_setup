/*
routers.go
Author: Naveenraj O M
Description:  Sets up and returns the Fiber application with defined routes for handling SME Scroll.
*/

package router

import (
	"mcommerce/internal/controller"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func GetRouter() *fiber.App {
	app := fiber.New()

	//adding cors
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH",
	}))

	// middleware
	app.Use(logger.New())
	app.Use(recover.New())

	app.Get("/", controller.HealthCheckHandler)

	//API group for user
	user_api_v1 := app.Group("/app/v1")

	user_api_v1.Get("/language", controller.HealthCheckHandler)

	return app
}
