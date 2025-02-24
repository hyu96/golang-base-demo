package healthz

import (
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"

	httpserver "github.com/huydq/gokits/libs/transport/http"
)

type HealthZBiz struct {
	api *httpserver.HttpServer
}

func NewHealthZBiz(httpServer *httpserver.HttpServer) *HealthZBiz {

	gr := httpServer.App.Group(viper.GetString("Env.HealthZPrefix"))
	gr.Get(
		"/api/v2.0/healthz/status",
		func(c *fiber.Ctx) error {
			return httpserver.WriteSuccessEmptyContent(c)
		},
	)

	return &HealthZBiz{api: httpServer}
}
