package iconst

type AuthKey string

const (
	KHeader_XApiKeyData = "X-Api-Key-Data"

	KAuthTokenInfo AuthKey = "tokenInfo"

	KAuthUsername           AuthKey = "Username"
	KAuthDeviceKey          AuthKey = "DeviceKey"
	KAuthWorkingSiteId      AuthKey = "WorkingSiteId"
	KAuthWorkerSiteId       AuthKey = "WorkerSiteId"
	KAuthWorkingSiteIntId   AuthKey = "WorkingSiteIntId"
	KAuthWorkerSiteIntId    AuthKey = "WorkerSiteIntId"
	KAuthUserID             AuthKey = "UserID"
	KAuthDisplayName        AuthKey = "DisplayName"
	KAuthMerchantLayerLevel AuthKey = "MerchantLayerLevel"
	KAuthMerchantLayerId    AuthKey = "MerchantLayerId"
)

func (a AuthKey) String() string {
	return string(a)
}
