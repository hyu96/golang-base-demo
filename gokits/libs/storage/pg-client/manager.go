package csql

import (
	"fmt"
	"sync"

	"github.com/huydq/gokits/libs/ilog"
)

type sqlClientManager struct {
	sqlClients sync.Map
}

var sqlClientsManagerInstance = &sqlClientManager{}

// default value env key is "MySQL";
// if configKeys was set, key env will be first value (not empty) of this;
func InstallSQLClientManager(configKeys ...string) {
	sqlConfigs := getConfigFromEnv(configKeys...)

	for _, config := range sqlConfigs {

		client := NewSqlxDB(config)
		if client == nil {
			err := fmt.Errorf("InstallSQLClientsManager - NewSqlxDB {%v} error", config)
			ilog.Errorf("InstallSQLClientsManager - Error: %v", err)

			panic(err)
		}

		if config.Name == "" {
			err := fmt.Errorf("InstallSQLClientsManager - config error: config.Name is empty")
			ilog.Errorf("InstallSQLClientsManager - Error: %v", err)

			panic(err)
		}
		if val, ok := sqlClientsManagerInstance.sqlClients.Load(config.Name); ok {
			err := fmt.Errorf("InstallSQLClientsManager - config error: duplicated config.Name {%v}", val)
			ilog.Errorf("InstallSQLClientsManager - Error: %v", err)

			panic(err)
		}

		sqlClientsManagerInstance.sqlClients.Store(config.Name, client)
	}
}

func GetSQLClient(dbName string) (client *SQLClient) {
	if val, ok := sqlClientsManagerInstance.sqlClients.Load(dbName); ok {
		if client, ok = val.(*SQLClient); ok {
			return
		}
	}

	ilog.Infof("GetSQLClient - Not found client: %s", dbName)
	return
}

func GetSQLClientManager() *sync.Map {
	return &sqlClientsManagerInstance.sqlClients
}
