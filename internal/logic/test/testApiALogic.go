package test

import (
	"context"

	"github.com/pz2147/p-api-1/internal/svc"
	"github.com/pz2147/p-api-1/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TestApiALogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTestApiALogic(ctx context.Context, svcCtx *svc.ServiceContext) TestApiALogic {
	return TestApiALogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TestApiALogic) TestApiA(req types.PingPeq) (resp *types.PingRes, err error) {


	return
}
