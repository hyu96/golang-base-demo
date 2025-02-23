package itimer

import (
	"fmt"
	"strings"
	"time"
)

const (
	TimestampFormat = `2006-01-02T15:04:05.999Z07:00`
)

//
// use carbon (https://github.com/golang-module/carbon) khi có thể
//

// Try parse date to time
// 2006-1
// 2006-01
// 2006-1-2
// 2006-01-02
func TryParseDate(date string) (time.Time, error) {
	date = strings.Split(date, " ")[0]

	if t, e := time.Parse("2006-1", date); e == nil {
		return t, nil
	}

	if t, e := time.Parse("2006-01", date); e == nil {
		return t, nil
	}

	if t, e := time.Parse("2006-1-2", date); e == nil {
		return t, nil
	}

	if t, e := time.Parse("2006-01-02", date); e == nil {
		return t, nil
	}

	return time.Time{}, fmt.Errorf("can not parse this format: %s", date)
}

func GetTimestampData() string {
	t := time.Now()
	location, err := time.LoadLocation("Asia/Ho_Chi_Minh")
	if err != nil {
		return t.Format(TimestampFormat)
	}
	return t.In(location).Format(TimestampFormat)
}
