package http_server

import (
	"flag"
	"github.com/carlmjohnson/versioninfo"
	"github.com/huydq/gokits/app"
	"github.com/huydq/gokits/libs/client/grpc"
	csql "github.com/huydq/gokits/libs/storage/pg-client"
	credis "github.com/huydq/gokits/libs/storage/redis"
	httpserver "github.com/huydq/gokits/libs/transport/http"
	"github.com/huydq/order-service/internal/adapter/persistence/postgres"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// flagconf is the config flag.
	flagconf string
)

func init() {
	flag.StringVar(&flagconf, "conf", "configs/public.yaml", "config path, eg: -conf config.yaml")
}

func StartHttpServer() {
	versioninfo.AddFlag(nil)

	flag.Parse()

	app.InitServer(flagconf) // need to do at first

	httpServer := httpserver.NewHttpServer()

	csql.InstallSQLClientManager()
	credis.InstallRedisClientManager()
	grpc.InstallRPCClient()

	orderPgClient := postgres.NewOrderPostgresClient(csql.NewBasePostgresSqlxDB(csql.DB_ORDER_SERVICE))
	if orderPgClient == nil {
		panic("Get postgres client failed")
	}
	s, err := wireApp(httpServer, *orderPgClient)
	if err != nil {
		panic(err)
	}

	app.DoInstance(s)
}
