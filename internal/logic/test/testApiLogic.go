package test

import (
	"context"
	"errors"
	"github.com/pz2147/p-rpc-1/prpc1client"

	"github.com/pz2147/p-api-1/internal/svc"
	"github.com/pz2147/p-api-1/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
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

	logx.Infof("[TestApi] TestApi start")

	test1, rpcErr := l.svcCtx.PRpc1.Test1(l.ctx, &prpc1client.Test1Req{
		RInt: 1111,
		RStr: "2222",
		RMap: &prpc1client.Test1CellModel{
			Cp: "sdsd",
		},
	})
	if rpcErr != nil {
		logx.Errorf("[TestApi] 错误 %s", rpcErr)
		return nil, err
	} else if test1 == nil {
		logx.Errorf("[TestApi] 返回为空")
		return nil, errors.New("返回为空")
	}
	return &types.PingRes{
		Code: 0,
	}, nil
}
