package csql

import "github.com/jmoiron/sqlx"

func In(query string, args ...interface{}) (string, []interface{}, error) {
	return sqlx.In(query, args...)
}
