package test

import (
	"context"
	"errors"
	"github.com/pz2147/p-rpc-1/prpc1client"

	"github.com/pz2147/p-api-1/internal/svc"
	"github.com/pz2147/p-api-1/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
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

	logx.Infof("[TestApiA] TestApiA start")

	test2, rpcErr := l.svcCtx.PRpc1.Test2(l.ctx, &prpc1client.Test2Req{
		RInt: 1111,
		RStr: "2222",
	})
	if rpcErr != nil {
		logx.Errorf("[TestApi] 错误 %s", rpcErr)
		return nil, err
	} else if test2 == nil {
		logx.Errorf("[TestApi] 返回为空")
		return nil, errors.New("返回为空")
	}
	return &types.PingRes{
		Code: 0,
	}, nil
}
