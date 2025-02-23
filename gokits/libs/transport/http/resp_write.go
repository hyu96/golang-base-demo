package httpserver

import (
	"github.com/gofiber/fiber/v2"
	iconst "github.com/huydq/gokits/constants"
	"github.com/huydq/gokits/libs/ilog"
	com "github.com/huydq/proto/gen-go/common"
)

const (
	KSuccess = "success"
	KError   = "error"
)

type Response struct {
	Status       string      `json:"status,omitempty"`
	Code         int         `json:"code,omitempty"`
	Description  string      `json:"description,omitempty"`
	ErrorNumber  int32       `json:"error_number,omitempty"`
	ErrorCode    string      `json:"error_code,omitempty"`
	ErrorMessage string      `json:"error_message,omitempty"`
	Message      string      `json:"message,omitempty"`
	Data         interface{} `json:"data,omitempty"`
}

func WriteSuccessEmptyContent(c *fiber.Ctx, message ...string) error {
	res := Response{
		Status: KSuccess,
		Code:   200,
	}

	if len(message) > 0 {
		res.Message = message[0]
	}

	return c.JSON(res)
}

func WriteSuccess(c *fiber.Ctx, v interface{}) error {
	res := Response{
		Data:   v,
		Status: KSuccess,
		Code:   200,
	}

	ilog.Debugw(
		"Response Success",
		"Route", c.Request().URI().String(),
		"RequestID", c.UserContext().Value(iconst.KContextKeyRequestID).(string),
		"Response", res,
	)

	return c.JSON(res)
}

func WriteError(c *fiber.Ctx, err *com.ApiError) error {
	res := Response{
		Status:       KError,
		Code:         int(err.GetCode()),
		ErrorNumber:  err.GetErrorCode(),
		ErrorCode:    com.KErrorCode(err.GetErrorCode()).String(),
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
