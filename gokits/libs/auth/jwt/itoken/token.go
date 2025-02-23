package itoken

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

const (
	KTokenAuthorizedKey  string = "authorized"
	KTokenAccessUUIDKey  string = "access_uuid"
	KTokenRefreshUUIDKey string = "refresh_uuid"
	KTokenUserIDKey      string = "user_id"
	KTokenExpKey         string = "exp"
	KTokenSecureExpKey   string = "secure_exp"
	KTokenAvatarURLKey   string = "avatarUrl"
	KTokenUserNameKey    string = "userName"
	KTokenDeviceKindKey  string = "device_kind"
	KTokenDeviceIPKey    string = "device_ip"
)

const (
	KCacheExpiresInOneHour  = time.Hour
	KCacheExpiresInOneDay   = 24 * KCacheExpiresInOneHour
	KCacheExpiresInOneWeek  = 7 * KCacheExpiresInOneDay
	KCacheExpiresInOneMonth = 30 * KCacheExpiresInOneDay
)

type Token struct {
	UserID             string        `json:"user_id,omitempty"`
	UserName           string        `json:"username,omitempty"`
	DeviceKind         int           `json:"device_kind,omitempty"`
	AccessUUID         uuid.UUID     `json:"access_uuid,omitempty"`
	RefreshUUID        uuid.UUID     `json:"refresh_uuid,omitempty"`
	AvatarUrl          string        `json:"avatar_url,omitempty"`
	AccessToken        *jwt.Token    `json:"access_token,omitempty"`
	RefreshToken       *jwt.Token    `json:"refresh_token,omitempty"`
	SignedAccessToken  string        `json:"signed_access_token,omitempty"`
	SignedRefreshToken string        `json:"signed_refresh_token,omitempty"`
	AtExpires          time.Duration `json:"at_expires,omitempty"`
	RtExpires          time.Duration `json:"rt_expires,omitempty"`
	SecureExpire       time.Duration `json:"secure_expire,omitempty"`
}

func Sign(token *Token, signer string) (string, string, error) {
	signedAt, err := token.AccessToken.SignedString([]byte(signer))
	if err != nil {
		return "", "", err
	}
	token.SignedAccessToken = signedAt

	signedRt, err := token.RefreshToken.SignedString([]byte(signer))
	if err != nil {
		return "", "", err
	}
	token.SignedRefreshToken = signedRt
	return signedAt, signedRt, nil
}
