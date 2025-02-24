package credis

import (
	"context"
	"fmt"
	"sync"

	"github.com/huydq/gokits/libs/ilog"
)

type redisClientManager struct {
	mapRedisClients sync.Map
}

var clients = &redisClientManager{}

// default value env key is "Redis";
// if configKeys was set, key env will be first value (not empty) of this;
// For the config, if addr is not empty, create redis node connection,
// if addrs and master name is not empty, create redis sentinel connection,
// otherwise create redis cluster connection
func InstallRedisClientManager(configKeys ...string) {
	configs := getRedisConfigFromEnv(configKeys...)
	if len(configs) == 0 {
		err := fmt.Errorf("not found config for redis cluster manager")
		ilog.Errorf("InstallRedisClientManager - Error: %+v", err)

		panic(err)
	}

	for _, config := range configs {
		pool := NewRedisUniversalClient(config)
		clients.mapRedisClients.Store(config.Name, pool)
		status := pool.client.Ping(context.Background())
		if status.Err() != nil {
			panic(fmt.Errorf("redis [%s], err:[%s]", config.Name, status.Err().Error()))
		}

		ilog.Infof("[=]Redis: [name]: %s, %s!", config.Name, status.String())
	}
}

// For the config, if addr is not empty, create redis node connection,
// if addrs and master name is not empty, create redis sentinel connection,
// otherwise create redis cluster connection
func InstallRedisClientManagerWithConfig(configs []*RedisConfig) {
	if len(configs) == 0 {
		err := fmt.Errorf("not found config for redis cluster manager")
		panic(err)
	}

	for _, config := range configs {
		pool := NewRedisUniversalClient(config)
		clients.mapRedisClients.Store(config.Name, pool)
		status := pool.client.Ping(context.Background())
		if status.Err() != nil {
			panic(fmt.Errorf("redis [%s], err:[%s]", config.Name, status.Err().Error()))
		}

		ilog.Infof("[=]Redis: [name]: %s, %s!", config.Name, status.String())

	}
}

func GetRedisClient(redisName string) (client *RedisPool) {
	if val, ok := clients.mapRedisClients.Load(redisName); ok {
		if client, ok = val.(*RedisPool); ok {
			return
		}
	}

	ilog.Infof("GetRedisClient - Not found client: %s", redisName)
	return
}

func GetRedisClientManager() *sync.Map {
	return &clients.mapRedisClients
}
