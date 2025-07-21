package logic

import (
	"context"

	"ws_trading/api/internal/svc"
	"ws_trading/api/internal/types"
	"ws_trading/rpc/trading"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetTradingSignalLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSetTradingSignalLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetTradingSignalLogic {
	return &SetTradingSignalLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SetTradingSignalLogic) SetTradingSignal(req *types.SetTradingSignalRequest) (resp *types.Base, err error) {
	l.Logger.Infof("SetTradingSignalLogic %v", req)
	resp = new(types.Base)

	res, err := l.svcCtx.TradingService.SetTradingSignal(l.ctx, &trading.SetTradingSignalRequest{
		Signal: req.Signal,
	})
	if err != nil {
		l.Logger.Errorf("SetTradingSignalLogic Error %v", err)
		resp.Code = 1
		resp.Message = err.Error()
	}
	resp.Code = int(res.Code)
	resp.Message = res.Message

	return
}
