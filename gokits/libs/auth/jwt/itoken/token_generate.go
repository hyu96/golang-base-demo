package itoken

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"

	iconst "github.com/huydq/gokits/constants"
)

func Generate(accountID, username, avatar string, deviceKind int, deviceIP string, configs ...*JWTConfig) (*Token, error) {
	if conf == nil {
		if len(configs) == 0 {
			GetJWTConfig()
		} else {
			conf = configs[0]
		}
	}

	td := &Token{}
	td.UserID = accountID
	td.UserName = username
	td.DeviceKind = deviceKind

	if deviceKind == iconst.KDeviceKindMobile {
		td.AtExpires = KCacheExpiresInOneDay
	} else {
		td.AtExpires = KCacheExpiresInOneMonth
	}
	td.SecureExpire = KCacheExpiresInOneHour
	td.RtExpires = KCacheExpiresInOneWeek

	accessUUID := uuid.New()
	td.AccessUUID = accessUUID
	atClaims := jwt.MapClaims{}
	atClaims[KTokenAuthorizedKey] = true
	atClaims[KTokenAccessUUIDKey] = accessUUID.String()
	atClaims[KTokenUserIDKey] = accountID
	atClaims[KTokenExpKey] = time.Now().Add(td.AtExpires).Unix()
	atClaims[KTokenAvatarURLKey] = td.AvatarUrl
	atClaims[KTokenUserNameKey] = td.UserName
	atClaims[KTokenDeviceKindKey] = deviceKind
	atClaims[KTokenDeviceIPKey] = deviceIP
	atClaims[KTokenSecureExpKey] = time.Now().Add(td.SecureExpire).Unix()
	td.AccessToken = jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)

	refreshUUID := uuid.New()
	td.RefreshUUID = refreshUUID
	rtClaims := jwt.MapClaims{}
	rtClaims[KTokenRefreshUUIDKey] = refreshUUID.String()
	rtClaims[KTokenUserIDKey] = accountID
	rtClaims[KTokenExpKey] = time.Now().Add(td.RtExpires).Unix()
	rtClaims[KTokenAvatarURLKey] = td.AvatarUrl
	rtClaims[KTokenUserNameKey] = td.UserName
	rtClaims[KTokenDeviceKindKey] = deviceKind
	rtClaims[KTokenDeviceIPKey] = deviceIP
	td.RefreshToken = jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)

	_, _, err := Sign(td, conf.SecretKey)

	return td, err
}
