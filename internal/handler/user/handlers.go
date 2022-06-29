package handler

import (
	"github.com/pz2147/p-api-1/internal/common/resp"
	"net/http"

	"github.com/pz2147/p-api-1/internal/logic/user"
	"github.com/pz2147/p-api-1/internal/svc"
	"github.com/pz2147/p-api-1/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func UserLoginHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginReq
		if err := httpx.Parse(r, &req); err != nil {
			resp.Error(w, err)
			return
		}

		l := user.NewUserLoginLogic(r.Context(), ctx)
		res, err := l.UserLogin(req)
		if err != nil {
			resp.Error(w, err)
		} else {
			resp.OkJson(w, res)
		}
	}
}

func UserLogoutHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LogoutReq
		if err := httpx.Parse(r, &req); err != nil {
			resp.Error(w, err)
			return
		}

		l := user.NewUserLogoutLogic(r.Context(), ctx)
		res, err := l.UserLogout(req)
		if err != nil {
			resp.Error(w, err)
		} else {
			resp.OkJson(w, res)
		}
	}
}

func UserInfoHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserInfoReq
		if err := httpx.Parse(r, &req); err != nil {
			resp.Error(w, err)
			return
		}

		l := user.NewUserInfoLogic(r.Context(), ctx)
		res, err := l.UserInfo(req)
		if err != nil {
			resp.Error(w, err)
		} else {
			resp.OkJson(w, res)
		}
	}
}
