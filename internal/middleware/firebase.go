package middleware

import (
	"context"
	"fmt"
	"log"
	"mcommerce/internal/firebase"
	"mcommerce/internal/response"
	"net/http"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

func FirebaseMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx := context.Background()
		client, err := firebase.FirebaseApp.Auth(ctx)
		if err != nil {
			log.Printf("error getting Auth client: %v\n", err)
			return c.Status(http.StatusInternalServerError).JSON(response.ErrorResponse{
				ApiPath:      c.OriginalURL(),
				ErrorCode:    http.StatusInternalServerError,
				ErrorMessage: "Context error",
				ErrorTime:    time.Now(),
			})
		}

		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(http.StatusUnauthorized).JSON(response.ErrorResponse{
				ApiPath:      c.OriginalURL(),
				ErrorCode:    http.StatusUnauthorized,
				ErrorMessage: "Authorization header missing",
				ErrorTime:    time.Now(),
			})
		}

		idToken := strings.Replace(authHeader, "Bearer ", "", 1)
		token, err := client.VerifyIDToken(ctx, idToken)
		if err != nil {
			log.Printf("error verifying ID token: %v\n", err)
			return c.Status(http.StatusUnauthorized).JSON(response.ErrorResponse{
				ApiPath:      c.OriginalURL(),
				ErrorCode:    http.StatusUnauthorized,
				ErrorMessage: "Invalid token",
				ErrorTime:    time.Now(),
			})
		}

		fmt.Printf("claims %v", token.Claims)

		// Extract firebase email
		fireBaseEmail, isOk := token.Claims["email"].(string)
		if isOk {
			c.Locals("fireBaseEmail", fireBaseEmail)
		} else {
			return c.Status(http.StatusInternalServerError).JSON(response.ErrorResponse{
				ApiPath:      c.OriginalURL(),
				ErrorCode:    http.StatusInternalServerError,
				ErrorMessage: "Email not found in firebase",
				ErrorTime:    time.Now(),
			})
		}

		// Extract firebase name
		firebaseName, isOk := token.Claims["name"].(string)
		if isOk {
			c.Locals("firebaseName", firebaseName)
		} else {
			return c.Status(http.StatusInternalServerError).JSON(response.ErrorResponse{
				ApiPath:      c.OriginalURL(),
				ErrorCode:    http.StatusInternalServerError,
				ErrorMessage: "Name not found in firebase",
				ErrorTime:    time.Now(),
			})
		}

		log.Printf("Verified ID token: %v\n", token)

		return c.Next()
	}
}
