package svc

import (
	"github.com/pz2147/p-api-1/internal/config"
	"github.com/pz2147/p-api-1/internal/middleware"
	"github.com/pz2147/p-rpc-1/prpc1client"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config  config.Config
	PRpc1   prpc1client.PRpc1
	Gateway rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		PRpc1:   prpc1client.NewPRpc1(zrpc.MustNewClient(c.PRpc1)),
		Gateway: middleware.NewGatewayMiddleware().Handle,
	}
}
