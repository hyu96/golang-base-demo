package httpserver

import (
	"github.com/gofiber/fiber/v2"
	"github.com/huydq/gokits/libs/ilog"
	"github.com/huydq/gokits/libs/utilities/ijson"
	errormsg "github.com/huydq/proto/error-message"
	comm "github.com/huydq/proto/gen-go/common/v2"
)

type ReqParam struct {
	Param        string
	Name         string // Name of param when it convert to body
	DefaultValue string
}

func GetContextDataString(ctx *fiber.Ctx, key string, defaultValues ...string) string {
	defaultValue := ""
	if len(defaultValues) > 0 {
		defaultValue = defaultValues[0]
	}

	userUUIDRaw := ctx.Locals(key)
	if userUUIDRaw != nil {
		if res, ok := userUUIDRaw.(string); ok {
			return res
		}
	}

	return defaultValue
}

// Parse data from param and set all into body
func ParamToBody(reqParams ...ReqParam) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		if len(reqParams) == 0 {
			return c.Next()
		}

		data := make(map[string]interface{})
		if err := c.BodyParser(&data); err != nil {
			ilog.Errorf("request[%s] error invalid body = %s", string(c.Context().URI().FullURI()), err.Error())
			return WriteErrorV2(c, errormsg.NewApiError(c.UserContext(), comm.Code_ERR_INTERNAL_SERVER))
		}

		for _, reqParam := range reqParams {
			data[reqParam.Name] = c.Params(reqParam.Param, reqParam.DefaultValue)
		}

		c.Request().SetBody(ijson.ToJsonByte(data))

		return c.Next()
	}
}

// Parse data from param and set all into body
func QueryToBody(reqParams ...ReqParam) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		if len(reqParams) == 0 {
			return c.Next()
		}

		data := make(map[string]interface{})
		if err := c.BodyParser(&data); err != nil {
			ilog.Errorf("request[%s] error invalid body = %s", string(c.Context().URI().FullURI()), err.Error())
			return WriteErrorV2(c, errormsg.NewApiError(c.UserContext(), comm.Code_ERR_INTERNAL_SERVER))
		}

		for _, reqParam := range reqParams {
			data[reqParam.Name] = c.Query(reqParam.Param, reqParam.DefaultValue)
		}

		c.Request().SetBody(ijson.ToJsonByte(data))

		return c.Next()
	}
}

// Parse data from param and set all into context
func ParamToContex(reqParams ...ReqParam) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		if len(reqParams) == 0 {
			return c.Next()
		}

		for _, reqParam := range reqParams {
			c.Locals(reqParam.Name, c.Params(reqParam.Param, reqParam.DefaultValue))
		}

		return c.Next()
	}
}

// Parse data from query and set all into context
func QueryToContext(reqParams ...ReqParam) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		if len(reqParams) == 0 {
			return c.Next()
		}

		for _, reqParam := range reqParams {
			c.Locals(reqParam.Name, c.Query(reqParam.Param, reqParam.DefaultValue))
		}

		return c.Next()
	}
}

// Parse all data from body and set all into context
func BodyToContext(c *fiber.Ctx) error {
	data := make(map[string]interface{})
	if err := c.BodyParser(&data); err != nil {
		ilog.Errorf("request[%s] error invalid body = %s", string(c.Context().URI().FullURI()), err.Error())
		return WriteErrorV2(c, errormsg.NewApiError(c.UserContext(), comm.Code_ERR_INTERNAL_SERVER))
	}

	for k, v := range data {
		c.Locals(k, v)
	}

	return c.Next()
}

// Set data{Name, DefaultValue} to context
func DataToContext(reqParams ...ReqParam) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		if len(reqParams) == 0 {
			return c.Next()
		}

		for _, reqParam := range reqParams {
			c.Locals(reqParam.Name, reqParam.DefaultValue)
		}

		return c.Next()
	}
}
