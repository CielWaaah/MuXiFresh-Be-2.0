package logic

import (
	"context"

	"MuXiFresh-Be-2.0/app/task/cmd/rpc/assignment/internal/svc"
	"MuXiFresh-Be-2.0/app/task/cmd/rpc/assignment/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAssignmentInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAssignmentInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAssignmentInfoLogic {
	return &GetAssignmentInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAssignmentInfoLogic) GetAssignmentInfo(in *pb.GetAssignmentInfoReq) (*pb.GetAssignmentInfoResp, error) {
	assignment, err := l.svcCtx.AssignmentModelClient.FindOne(l.ctx, in.AssignmentID)
	if err != nil {
		return nil, err
	}
	return &pb.GetAssignmentInfoResp{
		TitleText: assignment.TitleText,
		Content:   assignment.Content,
		Urls:      assignment.Urls,
	}, nil
}
