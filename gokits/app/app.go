package app

import (
	"bytes"
	"fmt"
	"os"
	"os/signal"
	"runtime/debug"
	"syscall"
	"time"

	"github.com/spf13/viper"

	"github.com/huydq/gokits/libs/env"
	"github.com/huydq/gokits/libs/ilog"
)

type ServerInstance interface {
	Start()
	Stop()
}

// *** Require load before “DoInstance“
func InitServer(configPath string) {
	// get from env
	var err error
	configStr := os.Getenv("APP_CONFIG")

	if configPath != "" {
		fmt.Println("trying to resolve config file Path: ", configPath)
		viper.SetConfigFile(configPath)
		err = viper.ReadInConfig()
	} else if configStr != "" {
		fmt.Println("trying to resolve APP_CONFIG: ", configStr)
		viper.SetConfigType("json")
		err = viper.ReadConfig(bytes.NewBufferString(configStr))
	}

	if err != nil {
		panic(fmt.Errorf("==> Error load config:[%s]", err))
	}

	// Setup config to component
	env.SetupEnvConfig()

	// Init log
	ilog.InitLogger()
}

var ch = make(chan os.Signal, 1)

// MAIN
func DoInstance(instance ServerInstance) {
	if instance == nil {
		panic("instance is nil")
	}

	// Print Server Info
	if ilog.Config().Development {
		ilog.Info("==> Service RUN on [DEVELOPMENT] mode")
	} else {
		ilog.Info("==> Service RUN on [PRODUCTION] mode")
	}
	ilog.Infof("[-] Environment:%s", env.Config().Environment)
	ilog.Infof("[-] DCName:%s", env.Config().DCName)
	ilog.Infof("[-] PodName:%s", env.Config().PodName)
	ilog.Infof("[-] PodID:%s", env.Config().PodID)
	ilog.Infof("[-] ServiceName:%s", env.Config().ServiceName)

	// Recovery error handler
	defer func() {
		if f := recover(); f != nil {
			if err, ok := f.(error); ok {
				ilog.Errorf("Global App Recover Error: %+v", err)
				debug.PrintStack()
			}
		}
	}()

	// Step1: Server Start
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL, syscall.SIGHUP, syscall.SIGQUIT)
	go instance.Start()

	// Step2: Server Stop
	s2 := <-ch
	if i, ok := s2.(syscall.Signal); ok {
		ilog.Infof("instance recv os.Exit(%d) signal...", i)
	} else {
		ilog.Infof("instance exit... with code %d", i)
	}

	instance.Stop()

	ilog.Infof("instance quited! <====")

	time.Sleep(1 * time.Second)
	os.Exit(0)
}

func QuitInstance() {
	ch <- syscall.SIGQUIT
}
