/*
health_check.go
Author: Naveenraj O M
Description: Defines an API endpoint for health checks in the Fiber framework.
This endpoint responds with a simple "Hello world!" message to confirm that the application is up and running.

The HealthCheckHandler function is used to handle HTTP requests to the root path ("/") and send a basic response indicating that the service is operational.

Usage:
This handler is typically used to check the availability and responsiveness of the application in a production environment.
It is often utilized in health checks by load balancers and monitoring tools to ensure that the application is running smoothly.
*/
package controller

import "github.com/gofiber/fiber/v2"

func HealthCheckHandler(c *fiber.Ctx) error {
	// Send a "Hello world" response to the client
	c.SendString("Hello M-Commerce!")

	// Return nil as there are no errors to handle in this
	return nil
}
