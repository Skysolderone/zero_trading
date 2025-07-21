package logic

import (
	"context"

	"ws_trading/rpc/internal/svc"
	"ws_trading/rpc/trading"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTradingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetTradingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTradingLogic {
	return &GetTradingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetTradingLogic) GetTrading(in *trading.GetTradingRequest) (*trading.GetTradingResponse, error) {
	result := new(trading.GetTradingResponse)
	result.Code = 0
	result.Message = "success"
	return result, nil
}
