package main

import (
	"flag"
	"fmt"

	"ws_trading/rpc/internal/config"
	"ws_trading/rpc/internal/server"
	"ws_trading/rpc/internal/svc"
	"ws_trading/rpc/trading"
	"ws_trading/utils/nacos"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	configFile = flag.String("f", "etc/trading.yaml", "the config file")
	nacosFlag  = flag.Bool("b", false, "the nacos config")
)

func main() {
	flag.Parse()
	fmt.Println("trading start", *configFile)
	var c config.Config
	if *nacosFlag {
		conf.MustLoad(*configFile, &c)
		configBytes := nacos.NewNacosConfig("", "", "")
		err := conf.LoadFromYamlBytes(configBytes, &c)
		if err != nil {
			panic(err)
		}
	} else {
		conf.MustLoad(*configFile, &c)
	}
	// fmt.Println("trading config", conf.LoadConfig())
	fmt.Println("trading config", c.Etcd)
	ctx := svc.NewServiceContext(c)
	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		trading.RegisterTradingServiceServer(grpcServer, server.NewTradingServiceServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
