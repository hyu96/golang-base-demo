package grpc_server

import (
	"flag"
	"github.com/carlmjohnson/versioninfo"
	"github.com/huydq/gokits/app"
	serverRPC "github.com/huydq/gokits/libs/transport/grpc"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// flagconf is the config flag.
	flagconf string
)

func init() {
	flag.StringVar(&flagconf, "conf", "configs/public.yaml", "config path, eg: -conf config.yaml")
}

func StartGrpcServer() {
	versioninfo.AddFlag(nil)
	flag.Parse()
	app.InitServer(flagconf) // need to do at first

	rpcServer := serverRPC.NewRPCServer()
	s, err := wireApp(rpcServer)
	if err != nil {
		panic(err)
	}
	app.DoInstance(s)
}
