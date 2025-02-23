package credis

import (
	"fmt"
	"strings"
	"time"

	redis "github.com/go-redis/redis/v8"
	"github.com/huydq/gokits/libs/ilog"
)

type RedisPool struct {
	client redis.UniversalClient
	env    string

	Conf RedisConfig
}

func NewRedisUniversalClient(conf *RedisConfig) *RedisPool {
	if conf.Addr != "" {
		if conf.DialTimeout < 10 {
			conf.DialTimeout = 10
		}

		if conf.ReadTimeout < 10 {
			conf.ReadTimeout = 10
		}

		if conf.WriteTimeout < 10 {
			conf.WriteTimeout = 10
		}

		if conf.IdleTimeout < 60 {
			conf.IdleTimeout = 60
		}

		myConfig := &redis.Options{
			Addr:     conf.Addr,
			DB:       conf.DBNum,
			Username: conf.Username,
			Password: conf.Password,

			MaxRetries: conf.MaxRetries,

			DialTimeout:  time.Duration(conf.DialTimeout) * time.Second,
			ReadTimeout:  time.Duration(conf.ReadTimeout) * time.Second,
			WriteTimeout: time.Duration(conf.WriteTimeout) * time.Second,
			IdleTimeout:  time.Duration(conf.IdleTimeout) * time.Second,

			MinIdleConns: conf.Idle,
			PoolSize:     conf.Active,
		}

		myClient := redis.NewClient(myConfig)

		myEnv := fmt.Sprintf("[%s]tcp@%s", conf.Name, strings.Join(conf.Addrs, ";"))

		pool := &RedisPool{env: myEnv, client: myClient}
		return pool
	}

	if len(conf.Addrs) > 0 {
		if len(conf.MasterName) > 0 {
			myConfig := &redis.FailoverOptions{
				SentinelAddrs:    conf.Addrs,
				DB:               conf.DBNum,
				Username:         conf.Username,
				Password:         conf.Password,
				SentinelPassword: conf.Password,

				MaxRetries: conf.MaxRetries,

				DialTimeout:  time.Duration(conf.DialTimeout) * time.Second,
				ReadTimeout:  time.Duration(conf.ReadTimeout) * time.Second,
				WriteTimeout: time.Duration(conf.WriteTimeout) * time.Second,
				IdleTimeout:  time.Duration(conf.IdleTimeout) * time.Second,

				MinIdleConns: conf.Idle,
				PoolSize:     conf.Active,

				// Only cluster clients.
				RouteByLatency: true,

				// The sentinel master name.
				// Only failover clients.
				MasterName: conf.MasterName,
			}

			myClient := redis.NewFailoverClient(myConfig)

			myEnv := fmt.Sprintf("[%s]tcp@%s", conf.Name, strings.Join(conf.Addrs, ";"))

			pool := &RedisPool{env: myEnv, client: myClient}
			return pool
		}

		myConfig := &redis.ClusterOptions{
			Addrs:      conf.Addrs,
			Username:   conf.Username,
			Password:   conf.Password,
			MaxRetries: conf.MaxRetries,

			DialTimeout:  time.Duration(conf.DialTimeout) * time.Second,
			ReadTimeout:  time.Duration(conf.ReadTimeout) * time.Second,
			WriteTimeout: time.Duration(conf.WriteTimeout) * time.Second,
			IdleTimeout:  time.Duration(conf.IdleTimeout) * time.Second,

			MinIdleConns: conf.Idle,
			PoolSize:     conf.Active,

			// Only cluster clients.
			RouteByLatency: true,
		}

		myClient := redis.NewClusterClient(myConfig)

		myEnv := fmt.Sprintf("[%s]tcp@%s", conf.Name, strings.Join(conf.Addrs, ";"))

		pool := &RedisPool{env: myEnv, client: myClient}
		return pool
	}

	return nil
}

// Get func
func (p *RedisPool) Get() redis.UniversalClient {
	if p == nil || p.client == nil {
		ilog.Panicf("🛑 redis client nil!")
		panic("🛑 redis client nil!")
	}
	return p.client
}

// Close func
func (p *RedisPool) Close() error {
	return p.client.Close()
}
