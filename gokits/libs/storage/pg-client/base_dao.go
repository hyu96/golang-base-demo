package csql

import (
	"context"
	"database/sql"
	"sync"

	"github.com/huydq/gokits/libs/storage"
	"github.com/jmoiron/sqlx"
)

var mapBaseSqlClient = sync.Map{}

type BasePostgresSqlxDB struct {
	Client *sqlx.DB
}

type IBasePostgresSqlxDB interface {
	Transaction(ctx context.Context, fn func(ctx context.Context) error) (err error)
	TransactionWithOption(ctx context.Context, option *sql.TxOptions, fn func(ctx context.Context) error) (err error)
}

func NewBasePostgresSqlxDB(dbName string) BasePostgresSqlxDB {
	return storage.LoadClient(
		dbName,
		&mapBaseSqlClient,
		GetSQLClientManager(),
		func(client interface{}) BasePostgresSqlxDB {
			return BasePostgresSqlxDB{Client: client.(*sqlx.DB)}
		},
	)
}

type txContextKey struct{}

func (d *BasePostgresSqlxDB) Transaction(ctx context.Context, fn func(ctx context.Context) error) (err error) {
	tx, err := d.Client.BeginTxx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}
	var done bool

	defer func() {
		if !done {
			_ = tx.Rollback()
		}
	}()
	if err = fn(context.WithValue(ctx, txContextKey{}, tx)); err != nil {
		return err
	}
	done = true
	return tx.Commit()
}

func (d *BasePostgresSqlxDB) TransactionWithOption(ctx context.Context, option *sql.TxOptions, fn func(ctx context.Context) error) (err error) {
	tx, err := d.Client.BeginTxx(ctx, option)
	if err != nil {
		return err
	}
	var done bool

	defer func() {
		if !done {
			_ = tx.Rollback()
		}
	}()
	if err = fn(context.WithValue(ctx, txContextKey{}, tx)); err != nil {
		return err
	}
	done = true
	return tx.Commit()
}

func (d *BasePostgresSqlxDB) GetTx(ctx context.Context) ISqlClient {
	value := ctx.Value(txContextKey{})
	if value != nil {
		return value.(*sqlx.Tx)
	}
	return d.Client
}
