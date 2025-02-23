package httpserver

import (
	"github.com/gofiber/fiber/v2"
	comm "github.com/huydq/proto/gen-go/common/v2"

	iconst "github.com/huydq/gokits/constants"
	"github.com/huydq/gokits/libs/ilog"
)

func WriteSuccessV2(c *fiber.Ctx, v interface{}) error {
	res := Response{
		Data:   v,
		Status: KSuccess,
		Code:   200,
	}

	ilog.Infow(
		"Response Success",
		"Route", c.Request().URI().String(),
		"RequestID", c.UserContext().Value(iconst.KContextKeyRequestID).(string),
		"Response", res,
	)

	return c.JSON(res)
}

func WriteErrorV2(c *fiber.Ctx, err *comm.ApiError) error {
	res := Response{
		Status:       KError,
		Code:         int(err.GetCode()),
		ErrorNumber:  err.GetErrorCode(),
		ErrorCode:    comm.Code(err.GetErrorCode()).String(),
		ErrorMessage: err.GetMessage(),
	}

	ilog.Infow(
		"Response Error",
		"Route", c.Request().URI().String(),
		"RequestID", c.UserContext().Value(iconst.KContextKeyRequestID).(string),
		"Response", res,
	)

	return c.Status(int(err.Code)).JSON(res)
}
