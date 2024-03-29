// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	test "github.com/pz2147/p-api-1/internal/handler/test"
	user "github.com/pz2147/p-api-1/internal/handler/user"
	"github.com/pz2147/p-api-1/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/test",
				Handler: test.TestApiHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/test/a",
				Handler: test.TestApiAHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/test/b",
				Handler: test.TestApiBHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Gateway},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/user/login",
					Handler: user.UserLoginHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/user/logout",
					Handler: user.UserLogoutHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/user/:id",
					Handler: user.UserInfoHandler(serverCtx),
				},
			}...,
		),
	)
}
