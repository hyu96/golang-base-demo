package grpc

import (
	"context"
	"strings"

	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	iconst "github.com/huydq/gokits/constants"
	"github.com/huydq/gokits/libs/auth/jwt/itoken"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func rpcAuthFunc(ctx context.Context) (context.Context, error) {
	m := grpc.ServerTransportStreamFromContext(ctx)
	if strings.Contains(m.Method(), "ServerReflectionInfo") {
		return ctx, nil
	}

	token, err := grpc_auth.AuthFromMD(ctx, "Bearer")
	if err != nil {
		return nil, err
	}

	claims, _, err := itoken.ParseToken("Bearer " + token)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "invalid auth token: %v", err)
	}

	if _, ok := iconst.MapWorkingSiteID[claims.WorkingSiteId]; !ok {
		return nil, status.Errorf(codes.InvalidArgument, "[working_site_id] không hợp lệ")
	}

	grpc_ctxtags.Extract(ctx).Set(iconst.KAuthUsername.String(), claims.Username)
	grpc_ctxtags.Extract(ctx).Set(iconst.KAuthDeviceKey.String(), claims.DeviceKey)
	grpc_ctxtags.Extract(ctx).Set(iconst.KAuthWorkingSiteId.String(), claims.WorkingSiteId)
	grpc_ctxtags.Extract(ctx).Set(iconst.KAuthWorkerSiteId.String(), claims.WorkerSiteId)
	grpc_ctxtags.Extract(ctx).Set(iconst.KAuthUserID.String(), claims.Id)
	grpc_ctxtags.Extract(ctx).Set(iconst.KAuthDisplayName.String(), claims.Name)

	newCtx := context.WithValue(ctx, iconst.KAuthTokenInfo, claims)
	newCtx = context.WithValue(newCtx, iconst.KAuthWorkingSiteId, claims.WorkingSiteId)
	newCtx = context.WithValue(newCtx, iconst.KAuthWorkerSiteId, claims.WorkerSiteId)

	return newCtx, nil
}

func RpcAuthFunc(ctx context.Context) (context.Context, error) {
	return rpcAuthFunc(ctx)
}
