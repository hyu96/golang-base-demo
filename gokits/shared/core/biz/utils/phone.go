package utils

import (
	iconst "github.com/huydq/gokits/constants"
	"regexp"
)

func VerifyPhoneNumberFormat(phone string) (valid bool) {
	pattern := regexp.MustCompile(`^(?:0\d{9}|84\d{9})$`)
	return pattern.MatchString(phone)
}

func GetLocalPhone(phone string) string {
	if phone[:2] == iconst.InternationalPhonePrefix {
		return iconst.LocalPhonePrefix + phone[2:]
	}
	return phone
}
