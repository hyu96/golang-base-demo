package csql

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // import your used driver

	"github.com/huydq/gokits/libs/ilog"
)

type ISqlClient interface {
	sqlx.ExtContext
}

type SQLClient struct {
	*sqlx.DB
}

func (client *SQLClient) Get() *sqlx.DB {
	return client.DB
}

// NewSqlxDB type;
// For MySQL, posgreSQL
func NewSqlxDB(c *SQLConfig) *sqlx.DB {
	db, err := sqlx.Connect(c.Driver, c.DSN)
	if err != nil {
		panic(fmt.Sprintf("NewSqlxDB Connect db error: %s", err.Error()))
	}

	db.SetMaxOpenConns(c.MaxOpenConns)
	db.SetMaxIdleConns(c.MaxIdleConns)
	db.SetConnMaxLifetime(time.Duration(c.ConnMaxLifetime) * time.Second)
	db.SetConnMaxIdleTime(time.Duration(c.ConnMaxIdleTime) * time.Second)

	ilog.Infof("[=]NewSqlxDB: [dbname]: %s, [setting]: %+v", c.Name, db.Stats())

	return db
}
