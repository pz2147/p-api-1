package user

import (
	"context"
	"github.com/pz2147/p-api-1/common/ec"
	"github.com/pz2147/p-rpc-1/prpc1client"

	"github.com/pz2147/p-api-1/internal/svc"
	"github.com/pz2147/p-api-1/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
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

func (l *UserInfoLogic) UserInfo(req types.UserInfoReq) (resp *types.UserinfoResp, err error) {

	uid := req.Uid
	if len(uid) == 0 {
		return nil, ec.ParamInvalid
	}

	lResp, lErr := l.svcCtx.PRpc1.UserInfoGet(l.ctx, &prpc1client.UserInfoReq{
		Uid: uid,
	})
	if lErr != nil {
		l.Logger.Errorf("[UserLogin] 登录失败 %s", lErr.Error())
		return nil, lErr
	}

	return &types.UserinfoResp{
		Uid:      lResp.GetUid(),
		Nickname: lResp.GetNickname(),
		Pic:      lResp.GetPic(),
	}, nil
}
