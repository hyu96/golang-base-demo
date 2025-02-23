package httpserver

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	iconst "github.com/huydq/gokits/constants"
)

// New creates a new middleware handler
func NewGenRequestID() fiber.Handler {

	// Return new handler
	return func(c *fiber.Ctx) error {

		ctx := context.WithValue(c.UserContext(), iconst.KContextKeyRequestID, uuid.NewString())
		c.SetUserContext(ctx) // Add the request ID to user context

		// Continue stack
		return c.Next()
	}
}
