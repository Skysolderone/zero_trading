package svc

import (
	"ws_trading/api/internal/config"
)

type ServiceContext struct {
	Config config.Config
	// TradingService trading.TradingService
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
