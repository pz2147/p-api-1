package test

import (
	"context"

	"github.com/pz2147/p-api-1/internal/svc"
	"github.com/pz2147/p-api-1/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TestApiBLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTestApiBLogic(ctx context.Context, svcCtx *svc.ServiceContext) TestApiBLogic {
	return TestApiBLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TestApiBLogic) TestApiB(req types.PingPeq) (resp *types.PingRes, err error) {

	return
}
