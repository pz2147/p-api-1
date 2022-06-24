package test

import (
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
	"p-api-1/internal/logic/test"
	"p-api-1/internal/svc"
	"p-api-1/internal/types"
)

func TestApiHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PingPeq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := test.NewTestApiLogic(r.Context(), ctx)
		resp, err := l.TestApi(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
