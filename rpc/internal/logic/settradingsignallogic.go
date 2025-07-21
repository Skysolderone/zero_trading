package logic

import (
	"context"

	"ws_trading/rpc/internal/svc"
	"ws_trading/rpc/trading"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetTradingSignalLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSetTradingSignalLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetTradingSignalLogic {
	return &SetTradingSignalLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SetTradingSignalLogic) SetTradingSignal(in *trading.SetTradingSignalRequest) (*trading.SetTradingSignalResponse, error) {
	l.Logger.Infof("rpc logic SetTradingSignalLogic %v", in)
	result := new(trading.SetTradingSignalResponse)
	result.Code = 0
	result.Message = "success"
	return result, nil
}
