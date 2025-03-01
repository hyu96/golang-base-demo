package order_handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/huydq/order-service/internal/core/domain/dto"
)

func (h *OrderHandler) CreateOrder(ctx *fiber.Ctx) error {
	var req dto.CreateOrderRequestDTO
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	orderRes, err := h.orderService.CreateOrder(ctx.UserContext(), req)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.JSON(orderRes)
}
