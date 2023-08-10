// Code generated by goctl. DO NOT EDIT.
// Source: test.proto

package server

import (
	"context"

	"MuXiFresh-Be-2.0/app/test/rpc/internal/logic"
	"MuXiFresh-Be-2.0/app/test/rpc/internal/svc"
	"MuXiFresh-Be-2.0/app/test/rpc/pb"
)

type TestClientServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedTestClientServer
}

func NewTestClientServer(svcCtx *svc.ServiceContext) *TestClientServer {
	return &TestClientServer{
		svcCtx: svcCtx,
	}
}

func (s *TestClientServer) UserTest(ctx context.Context, in *pb.TestReq) (*pb.TestResp, error) {
	l := logic.NewUserTestLogic(ctx, s.svcCtx)
	return l.UserTest(in)
}

func (s *TestClientServer) CheckTestResult(ctx context.Context, in *pb.TestInfoReq) (*pb.TestInfoResp, error) {
	l := logic.NewCheckTestResultLogic(ctx, s.svcCtx)
	return l.CheckTestResult(in)
}

func (s *TestClientServer) JudgeUserType(ctx context.Context, in *pb.UserTypeReq) (*pb.UserTypeResp, error) {
	l := logic.NewJudgeUserTypeLogic(ctx, s.svcCtx)
	return l.JudgeUserType(in)
}
