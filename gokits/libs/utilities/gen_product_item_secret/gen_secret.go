package gen_secret

import (
	"github.com/twharmon/gouid"
)

const (
	KGENCODE_LIMIT = 12
)

func GenSerial(prefix string) string {
	serial := gouid.String(KGENCODE_LIMIT, gouid.UpperCaseAlphaNum)
	return prefix + serial
}

func GenSecretCode(prefix string) string {
	secret := gouid.String(KGENCODE_LIMIT, gouid.UpperCaseAlphaNum)
	return prefix + secret
}

func GenFullCode(prefix string) (serial, secret string) {
	return GenSerial(prefix), GenSecretCode(prefix)
}
