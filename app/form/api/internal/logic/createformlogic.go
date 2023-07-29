package logic

import (
	"MuXiFresh-Be-2.0/app/form/rpc/entryformclient"
	externalModel "MuXiFresh-Be-2.0/app/userauth/model"
	"MuXiFresh-Be-2.0/common/ctxData"
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"

	"MuXiFresh-Be-2.0/app/form/api/internal/svc"
	"MuXiFresh-Be-2.0/app/form/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateFormLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateFormLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateFormLogic {
	return &CreateFormLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateFormLogic) CreateForm(req *types.CreateReq) (resp *types.CreateResp, err error) {
	userId := ctxData.GetUserIdFromCtx(l.ctx)
	CtResp, err := l.svcCtx.FormClient.CreateForm(l.ctx, &entryformclient.CreateReq{
		UserId:        userId,
		Avatar:        req.Avatar,
		Name:          req.Name,
		StuNumber:     req.StuNumber,
		School:        req.School,
		Major:         req.Major,
		Grade:         req.Grade,
		Gender:        req.Gender,
		Email:         req.Email,
		QQ:            req.QQ,
		Phone:         req.Phone,
		Group:         req.Group,
		Reason:        req.Reason,
		Knowledge:     req.Knowledge,
		SelfIntro:     req.SelfIntro,
		ExtraQuestion: req.ExtraQuestion,
	})
	u, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return nil, err
	}
	f, err := primitive.ObjectIDFromHex(CtResp.FormID)
	if err != nil {
		return nil, err
	}

	_, err = l.svcCtx.UserInfoModelClient.Update(l.ctx, &externalModel.UserInfo{
		ID:          u,
		EntryFormID: f,
		UpdateAt:    time.Now(),
	})
	if err != nil {
		return nil, err
	}
	return &types.CreateResp{
		Flag: true,
	}, nil
}
