// Code generated by goctl. DO NOT EDIT.
// Source: intro.proto

package server

import (
	"context"

	"MuXiFresh-Be-2.0/app/intro/rpc/internal/logic"
	"MuXiFresh-Be-2.0/app/intro/rpc/internal/svc"
	"MuXiFresh-Be-2.0/app/intro/rpc/pb"
)

type IntroClientServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedIntroClientServer
}

func NewIntroClientServer(svcCtx *svc.ServiceContext) *IntroClientServer {
	return &IntroClientServer{
		svcCtx: svcCtx,
	}
}

func (s *IntroClientServer) GetGroupIntro(ctx context.Context, in *pb.GroupIntroReq) (*pb.GroupIntroResp, error) {
	l := logic.NewGetGroupIntroLogic(ctx, s.svcCtx)
	return l.GetGroupIntro(in)
}

func (s *IntroClientServer) GetRecruitInfo(ctx context.Context, in *pb.RecruitInfoReq) (*pb.RecruitInfoResp, error) {
	l := logic.NewGetRecruitInfoLogic(ctx, s.svcCtx)
	return l.GetRecruitInfo(in)
}
