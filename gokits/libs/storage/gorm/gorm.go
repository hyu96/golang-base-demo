package gorm

import (
	"database/sql"

	"github.com/huydq/gokits/libs/ilog"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type GormClient struct {
	Db    *gorm.DB
	SqlDB *sql.DB
}

// NewRPCClient func
func NewGormClient(conf *GormConfig) *GormClient {
	db, err := gorm.Open(mysql.Open(conf.DSN), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	// Thiết lập số lượng kết nối trong pool
	sqlDB.SetMaxIdleConns(conf.Idle)
	sqlDB.SetMaxOpenConns(conf.Active)

	ilog.Infof("[-] Installed sqlgorm client done!")

	return &GormClient{
		Db:    db,
		SqlDB: sqlDB,
	}
}
