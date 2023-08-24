// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package server

import (
	"context"

	"MuXiFresh-Be-2.0/app/user/cmd/rpc/user/internal/logic"
	"MuXiFresh-Be-2.0/app/user/cmd/rpc/user/internal/svc"
	"MuXiFresh-Be-2.0/app/user/cmd/rpc/user/pb"
)

type UserClientServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedUserClientServer
}

func NewUserClientServer(svcCtx *svc.ServiceContext) *UserClientServer {
	return &UserClientServer{
		svcCtx: svcCtx,
	}
}

func (s *UserClientServer) GetUserInfo(ctx context.Context, in *pb.GetUserInfoReq) (*pb.GetUserInfoResp, error) {
	l := logic.NewGetUserInfoLogic(ctx, s.svcCtx)
	return l.GetUserInfo(in)
}

func (s *UserClientServer) SetUserInfo(ctx context.Context, in *pb.SetUserInfoReq) (*pb.SetUserInfoResp, error) {
	l := logic.NewSetUserInfoLogic(ctx, s.svcCtx)
	return l.SetUserInfo(in)
}

func (s *UserClientServer) SetUserType(ctx context.Context, in *pb.SetUserTypeReq) (*pb.SetUserTypeResp, error) {
	l := logic.NewSetUserTypeLogic(ctx, s.svcCtx)
	return l.SetUserType(in)
}

func (s *UserClientServer) GetAdminList(ctx context.Context, in *pb.GetAdminListReq) (*pb.GetAdminListResp, error) {
	l := logic.NewGetAdminListLogic(ctx, s.svcCtx)
	return l.GetAdminList(in)
}

func (s *UserClientServer) GetUserType(ctx context.Context, in *pb.GetUserTypeReq) (*pb.GetUserTypeResp, error) {
	l := logic.NewGetUserTypeLogic(ctx, s.svcCtx)
	return l.GetUserType(in)
}