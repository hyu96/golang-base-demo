package itoken

import (
	"context"
	"github.com/golang-jwt/jwt/v4"
)

const (
	RefreshToken  = "refresh"
	AccessToken   = "access"
	RootWorkspace = "ROOT"
)

type claimCtxKey struct{}

func ContextWithClaims(ctx context.Context, claims *Claims) context.Context {
	return context.WithValue(ctx, claimCtxKey{}, claims)
}

func ClaimsFromContext(ctx context.Context) (*Claims, bool) {
	claims, ok := ctx.Value(claimCtxKey{}).(*Claims)
	return claims, ok
}

//func MustClaimsFromContext(ctx context.Context) *Claims {
//	claims, ok := ClaimsFromContext(ctx)
//	if !ok {
//		panic(errors.Unauthorized("UNAUTHENTICATED", "unauthenticated"))
//	}
//	return claims
//}

type Claims struct {
	jwt.RegisteredClaims
	Identifier    string `json:"identifier"`
	Username      string `json:"username"`
	Name          string `json:"name"`
	Id            string `json:"id"`
	DeviceKey     string `json:"device_key"`
	WorkingSiteId string `json:"working_site_id"`
	WorkerSiteId  string `json:"worker_site_id"`
	Version       string `json:"version"`
	WorkspaceCode string `json:"workspaceCode"`
	IsService     bool   `json:"isService"`
	Type          string `json:"typ"`
}

func (c *Claims) String() string {
	if c.IsService {
		return c.Identifier + "@" + c.WorkspaceCode + ".service"
	}
	return c.Identifier + "@" + c.WorkspaceCode + "user"
}

type ApiKeyClaim struct {
	IsService          bool  `json:"is_service"`
	WorkingSiteId      int64 `json:"working_site_id"`
	WorkerSiteId       int64 `json:"worker_site_id"`
	MerchantLayerLevel int64 `json:"merchant_layer_level"`
	MerchantLayerId    int64 `json:"merchant_layer_id"`
}
