package httpserver

import (
	"fmt"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/huydq/gokits/libs/env"
)

type HttpServer struct {
	*fiber.App
	conf *HttpConfig
}

var (
	httpInstance *HttpServer
)

func NewHttpServer() *HttpServer {
	if httpInstance != nil {
		return httpInstance
	}

	setupHttpServerConfig()

	fiberApp := fiber.New()
	fiberApp.Use(cors.New())
	fiberApp.Use(recover.New(recover.Config{EnableStackTrace: true}))
	fiberApp.Use(NewGenRequestID())
	fiberApp.Use(logger.New(logger.Config{Next: func(c *fiber.Ctx) bool {
		return strings.Contains(c.Request().URI().String(), "healthz")
	}}))

	fiberApp.Get("/ping", ping)

	httpInstance = &HttpServer{
		App:  fiberApp,
		conf: Config,
	}

	return httpInstance
}

func GetGAPIInstance() *HttpServer {
	if httpInstance == nil {
		err := fmt.Errorf("you need call NewHttpServer api first")
		panic(err)
	}

	return httpInstance
}

func (api *HttpServer) GetConfig() *HttpConfig {
	return api.conf
}

// Need run in a goroutine
func (api *HttpServer) Serve() {
	if err := api.App.Listen(api.conf.Port); err != nil {
		_ = fmt.Errorf("HttpServer::Serve - Listen Error: %+v", err)
	}
}

func (api *HttpServer) Stop() {
	shutdownTimeout := time.Duration(env.Config().ShutdownTimeout) * time.Second
	if err := api.App.ShutdownWithTimeout(shutdownTimeout); err != nil {
		fmt.Printf("HttpServer::Stop - Error: %+v\n", err)
	}
}
