package logic

import (
	"MuXiFresh-Be-2.0/app/form/model"
	"MuXiFresh-Be-2.0/app/form/rpc/internal/svc"
	"MuXiFresh-Be-2.0/app/form/rpc/pb"
	"MuXiFresh-Be-2.0/common/xerr"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateFormLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateFormLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateFormLogic {
	return &UpdateFormLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateFormLogic) UpdateForm(in *pb.CreateReq) (*pb.CreateResp, error) {
	u, err := primitive.ObjectIDFromHex(in.UserId)
	if err != nil {
		return nil, xerr.ErrExistInvalidId.Status()
	}
	f, err := primitive.ObjectIDFromHex(in.FormId)
	if err != nil {
		return nil, xerr.ErrExistInvalidId.Status()
	}
	//form := model.EntryForm{}
	//copier.Copy(&form,in)
	updateRet, err := l.svcCtx.FormClient.Update(l.ctx, &model.EntryForm{
		UserId:        u,
		ID:            f,
		Avatar:        in.Avatar,
		Major:         in.Major,
		Grade:         in.Grade,
		Gender:        in.Gender,
		Phone:         in.Phone,
		Group:         in.Group,
		Reason:        in.Reason,
		Knowledge:     in.Knowledge,
		SelfIntro:     in.SelfIntro,
		ExtraQuestion: in.ExtraQuestion,
		UpdateAt:      time.Now(),
	})
	
	if err != nil {
		switch err {
		case model.ErrNotFound:
			return nil, xerr.ErrNotFind.Status()
		default:
			return nil, xerr.NewErrCode(xerr.DB_ERROR).Status()
		}
	}

	return &pb.CreateResp{
		FormID: fmt.Sprint(updateRet.UpsertedID),
	}, nil
}
