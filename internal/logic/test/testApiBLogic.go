package test

import (
	"context"
	"errors"
	"github.com/pz2147/p-rpc-1/prpc1client"

	"github.com/pz2147/p-api-1/internal/svc"
	"github.com/pz2147/p-api-1/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
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

	logx.Infof("[TestApiB] TestApiB start")

	test3, rpcErr := l.svcCtx.PRpc1.Test3(l.ctx, &prpc1client.Test3Req{
		RInt: 1111,
		RStr: "2222",
	})
	if rpcErr != nil {
		logx.Errorf("[TestApi] 错误 %s", rpcErr)
		return nil, err
	} else if test3 == nil {
		logx.Errorf("[TestApi] 返回为空")
		return nil, errors.New("返回为空")
	}
	return &types.PingRes{
		Code: 0,
	}, nil
}
