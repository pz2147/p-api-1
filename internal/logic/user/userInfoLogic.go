package user

import (
	"context"

	"github.com/pz2147/p-api-1/internal/svc"
	"github.com/pz2147/p-api-1/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) UserInfoLogic {
	return UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo(req types.UidReq) (resp *types.UserinfoResp, err error) {
	// todo: add your logic here and delete this line

	return
}
