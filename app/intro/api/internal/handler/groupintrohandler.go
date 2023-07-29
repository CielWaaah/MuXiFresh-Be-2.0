package handler

import (
	"MuxiFresh2.0/MuXiFresh-Be-2.0/app/intro/api/internal/logic"
	"MuxiFresh2.0/MuXiFresh-Be-2.0/app/intro/api/internal/svc"
	"MuxiFresh2.0/MuXiFresh-Be-2.0/app/intro/api/internal/types"
	"greet/response"
	"net/http"
)

func GroupIntroHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GroupIntroReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewGroupIntroLogic(r.Context(), svcCtx)
		resp, err := l.GroupIntro(&req)
		response.Response(w, resp, err)

	}
}
