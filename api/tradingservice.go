package main

import (
	"flag"
	"fmt"

	"ws_trading/api/internal/config"
	"ws_trading/api/internal/handler"
	"ws_trading/api/internal/svc"
	"ws_trading/utils/nacos"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/threading"
	"github.com/zeromicro/go-zero/rest"
)

var (
	configFile = flag.String("f", "etc/tradingservice.yaml", "the config file")
	nacosFlag  = flag.Bool("b", false, "the nacos config")
)

func main() {
	flag.Parse()
	fmt.Println("tradingservice start", *configFile)
	var c config.Config
	// 异常捕获
	defer threading.GoSafe(func() {})
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
	fmt.Println("tradingservice config", c.TradingService)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
