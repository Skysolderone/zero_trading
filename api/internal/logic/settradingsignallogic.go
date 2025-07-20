package logic

import (
	"context"

	"ws_trading/api/internal/svc"
	"ws_trading/api/internal/types"

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
	resp = new(types.Base)
	//编写rpc服务后调用
	// res, err := l.svcCtx.TradingService.SetTradingSignal(l.ctx, &trading.SetTradingSignalRequest{
	// 	Signal: req.Signal,
	// })
	// if err != nil {
	// 	resp.Code = 1
	// 	resp.Message = err.Error()
	// }
	resp.Code = 0
	resp.Message = "success"

	return
}
