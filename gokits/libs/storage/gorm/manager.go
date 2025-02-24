package gorm

import (
	"fmt"
	"sync"

	"github.com/huydq/gokits/libs/ilog"
)

var gormClients sync.Map

func InstallSqlGormClient() {
	LoadGormClientConfig()

	for _, config := range configs {
		client := NewGormClient(&config)
		if client == nil {
			panic(fmt.Errorf("InstallSqlGormClient - NewGormClient {%v} error", config))
		}

		if config.Name == "" {
			panic(fmt.Errorf("InstallSqlGormClient - config error: config.Name is empty"))
		}
		if val, ok := gormClients.Load(config.Name); ok {
			panic(fmt.Errorf("InstallSQLClientsManager - config error: duplicated config.Name {%v}", val))
		}

		gormClients.Store(config.Name, client)
	}
}

func GetGormClient(dbName string) (client *GormClient) {
	if val, ok := gormClients.Load(dbName); ok {
		if client, ok = val.(*GormClient); ok {
			return
		}
	}

	ilog.Infof("GetGormClient - Not found client: %s", dbName)
	return
}

func GetGormClientManager() *sync.Map {
	return &gormClients
}
