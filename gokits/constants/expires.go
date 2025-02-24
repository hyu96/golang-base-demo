package iconst

import "time"

// util for expire time
const (
	KExpiresInForever        = 0
	KExpiresInOneMinute      = time.Minute
	KExpiresInTenMinutes     = time.Minute * 10
	KExpiresInFifteenMinutes = time.Minute * 15
	KExpiresInOneHour        = time.Hour
	KExpiresInTwoHour        = KExpiresInOneHour * 2
	KExpiresInOneDay         = KExpiresInOneHour * 24
	KExpiresInThreeDay       = KExpiresInOneDay * 3
	KExpiresInOneMonth       = KExpiresInOneDay * 30
	KExpiresInThreeMonths    = KExpiresInOneMonth * 3
	KExpiresInSixMonths      = KExpiresInOneMonth * 6
	KExpiresInOneYear        = KExpiresInOneMonth * 12
)
