package order_handler

import (
	"github.com/gofiber/fiber/v2"
	pbOrderMgmt "github.com/huydq/proto/gen-go/order"
)

func (h *OrderHandler) CreateOrder(ctx *fiber.Ctx) error {
	var req pbOrderMgmt.CreateOrderRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	orderRes, err := h.orderService.CreateOrder(ctx.UserContext(), req)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.JSON(orderRes)
}
