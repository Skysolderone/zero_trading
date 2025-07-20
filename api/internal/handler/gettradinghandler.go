package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"ws_trading/api/internal/logic"
	"ws_trading/api/internal/svc"
	"ws_trading/api/internal/types"
)

func GetTradingHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetTradingRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewGetTradingLogic(r.Context(), svcCtx)
		resp, err := l.GetTrading(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
