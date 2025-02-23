package itimer

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"time"
)

const (
	TIME_FORMAT_DATE = "2006-01-02"
	TIME_FORMAT_TIME = "2006-01-02 15:04:05"
)

type SqlTimeDate time.Time

func (ct *SqlTimeDate) Scan(value interface{}) error {
	if bt, ok := value.([]byte); ok {
		t, err := time.Parse(TIME_FORMAT_DATE, string(bt))
		if err != nil {
			return err
		}
		*ct = SqlTimeDate(t)
		return nil
	}
	return errors.New("invalid time format")
}

func (t SqlTimeDate) Value() (driver.Value, error) {
	return driver.Value(time.Time(t).Format(TIME_FORMAT_DATE)), nil
}

func (t *SqlTimeDate) UnmarshalText(value string) error {
	dd, err := time.Parse(TIME_FORMAT_DATE, value)
	if err != nil {
		return err
	}
	*t = SqlTimeDate(dd)
	return nil
}

func (t SqlTime) SqlTimeDate() ([]byte, error) {
	//do your serializing here
	stamp := fmt.Sprintf("\"%s\"", time.Time(t).Format(TIME_FORMAT_DATE))
	return []byte(stamp), nil
}

type SqlTime time.Time

func (ct *SqlTime) Scan(value interface{}) error {
	if bt, ok := value.([]byte); ok {
		t, err := time.Parse(TIME_FORMAT_TIME, string(bt))
		if err != nil {
			return err
		}
		*ct = SqlTime(t)
		return nil
	}
	return errors.New("invalid time format")
}

func (t SqlTime) Value() (driver.Value, error) {
	return driver.Value(time.Time(t).Format(TIME_FORMAT_TIME)), nil
}

func (t *SqlTime) UnmarshalText(value string) error {
	dd, err := time.Parse(TIME_FORMAT_TIME, value)
	if err != nil {
		return err
	}
	*t = SqlTime(dd)
	return nil
}

func (t SqlTime) MarshalJSON() ([]byte, error) {
	//do your serializing here
	stamp := fmt.Sprintf("\"%s\"", time.Time(t).Format(TIME_FORMAT_TIME))
	return []byte(stamp), nil
}
