package svc

import (
	"ws_trading/api/internal/config"
	"ws_trading/rpc/tradingservice"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config         config.Config
	TradingService tradingservice.TradingService
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:         c,
		TradingService: tradingservice.NewTradingService(zrpc.MustNewClient(c.TradingService)),
	}
}
