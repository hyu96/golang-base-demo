package grpc

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/huydq/gokits/libs/env"
	hgrpc "github.com/huydq/gokits/libs/grpc-helper/middleware"
	"github.com/huydq/gokits/libs/ilog"
	com "github.com/huydq/proto/gen-go/common"
)

func BizUnaryRecoveryHandler2(ctx context.Context, p interface{}) (err error) {
	switch code := p.(type) {
	case *com.ApiError:
		md, _ := hgrpc.RPCErrorToMD(code)
		grpc.SetTrailer(ctx, md)
		err = code
	default:
		err = status.Errorf(codes.Unknown, "panic unknown triggered: %v", p)
		errDesc := fmt.Sprintf("💣💣💣 At %s.\nPanic unknown triggered: %+v", env.Config().Environment, err.Error())
		ilog.Errorf("BizUnaryRecoveryHandler2 - Error: %v", errDesc)
	}

	return
}
