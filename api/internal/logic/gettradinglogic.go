package logic

import (
	"context"

	"ws_trading/api/internal/svc"
	"ws_trading/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTradingLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTradingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTradingLogic {
	return &GetTradingLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTradingLogic) GetTrading(req *types.GetTradingRequest) (resp *types.BaseResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
