package dtm

import (
	"context"
	"net/url"

	"github.com/dtm-labs/logger"
	"github.com/gofiber/fiber/v2"
	"github.com/huydq/gokits/third_party/dtm-client/dtmcli"
	"github.com/huydq/gokits/third_party/dtm-client/dtmgrpc"
)

func MustBarrierFromFiber(c *fiber.Ctx) *dtmcli.BranchBarrier {
	queryString := c.Request().URI().QueryString()
	values, _ := url.ParseQuery(string(queryString))
	ti, err := dtmcli.BarrierFromQuery(values)
	logger.FatalIfError(err)
	return ti
}

func MustBarrierFromGrpc(ctx context.Context) *dtmcli.BranchBarrier {
	ti, err := dtmgrpc.BarrierFromGrpc(ctx)
	logger.FatalIfError(err)
	return ti
}
