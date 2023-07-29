package handler

import (
	"MuxiFresh2.0/MuXiFresh-Be-2.0/app/schedule/api/internal/logic"
	"MuxiFresh2.0/MuXiFresh-Be-2.0/app/schedule/api/internal/svc"
	"MuxiFresh2.0/MuXiFresh-Be-2.0/app/schedule/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
	"greet/response"
	"net/http"
)

func CreateScheduleHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewCreateScheduleLogic(r.Context(), svcCtx)
		resp, err := l.CreateSchedule(&req)
		response.Response(w, resp, err)

	}
}