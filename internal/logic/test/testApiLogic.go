package test

import (
	"context"

	"github.com/pz2147/p-api-1/internal/svc"
	"github.com/pz2147/p-api-1/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TestApiLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTestApiLogic(ctx context.Context, svcCtx *svc.ServiceContext) TestApiLogic {
	return TestApiLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TestApiLogic) TestApi(req types.PingPeq) (resp *types.PingRes, err error) {

	return
}
