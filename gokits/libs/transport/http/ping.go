package httpserver

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

func ping(c *fiber.Ctx) error {
	return c.JSON(map[string]interface{}{
		"status":  200,
		"message": "pong",
		"data": map[string]interface{}{
			"message": "pong",
			"time":    time.Now().UTC().Format("2006-01-02 15:04:05"),
			"status":  "oke",
			"from":    c.Context().RemoteAddr().String(),
			"agent":   string(c.Context().UserAgent()),
		},
	})
}
