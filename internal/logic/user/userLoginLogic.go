package user

import (
	"context"
	"errors"

	"github.com/pz2147/p-api-1/internal/svc"
	"github.com/pz2147/p-api-1/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type UserLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) UserLoginLogic {
	return UserLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLoginLogic) UserLogin(req types.LoginReq) (resp *types.LoginResp, err error) {

	phone := req.Phone
	password := req.Password
	if len(phone) == 0 || len(password) == 0 {
		return nil, errors.New("参数错误")
	}

	l.Logger.Infof("%s %s", phone, password)

	return &types.LoginResp{
		Uid:      "123456789",
		Nickname: phone,
		Pic:      "pic/pic/" + phone,
	}, nil
}
