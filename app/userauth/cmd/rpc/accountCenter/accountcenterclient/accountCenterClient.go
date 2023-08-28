// Code generated by goctl. DO NOT EDIT.
// Source: accountCenter.proto

package accountcenterclient

import (
	"context"

	"MuXiFresh-Be-2.0/app/userauth/cmd/rpc/accountCenter/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	CcnuLoginReq     = pb.CcnuLoginReq
	CcnuLoginResp    = pb.CcnuLoginResp
	LoginVerifyReq   = pb.LoginVerifyReq
	LoginVerifyResp  = pb.LoginVerifyResp
	RegisterDataReq  = pb.RegisterDataReq
	RegisterDataResp = pb.RegisterDataResp
	SetEmailReq      = pb.SetEmailReq
	SetEmailResp     = pb.SetEmailResp
	SetPasswordReq   = pb.SetPasswordReq
	SetPasswordResp  = pb.SetPasswordResp
	SetStudentIDReq  = pb.SetStudentIDReq
	SetStudentIDResp = pb.SetStudentIDResp

	AccountCenterClient interface {
		Register(ctx context.Context, in *RegisterDataReq, opts ...grpc.CallOption) (*RegisterDataResp, error)
		Login(ctx context.Context, in *LoginVerifyReq, opts ...grpc.CallOption) (*LoginVerifyResp, error)
		SetPassword(ctx context.Context, in *SetPasswordReq, opts ...grpc.CallOption) (*SetPasswordResp, error)
		CcnuLogin(ctx context.Context, in *CcnuLoginReq, opts ...grpc.CallOption) (*CcnuLoginResp, error)
		SetStudentID(ctx context.Context, in *SetStudentIDReq, opts ...grpc.CallOption) (*SetStudentIDResp, error)
		SetEmail(ctx context.Context, in *SetEmailReq, opts ...grpc.CallOption) (*SetEmailResp, error)
	}

	defaultAccountCenterClient struct {
		cli zrpc.Client
	}
)

func NewAccountCenterClient(cli zrpc.Client) AccountCenterClient {
	return &defaultAccountCenterClient{
		cli: cli,
	}
}

func (m *defaultAccountCenterClient) Register(ctx context.Context, in *RegisterDataReq, opts ...grpc.CallOption) (*RegisterDataResp, error) {
	client := pb.NewAccountCenterClientClient(m.cli.Conn())
	return client.Register(ctx, in, opts...)
}

func (m *defaultAccountCenterClient) Login(ctx context.Context, in *LoginVerifyReq, opts ...grpc.CallOption) (*LoginVerifyResp, error) {
	client := pb.NewAccountCenterClientClient(m.cli.Conn())
	return client.Login(ctx, in, opts...)
}

func (m *defaultAccountCenterClient) SetPassword(ctx context.Context, in *SetPasswordReq, opts ...grpc.CallOption) (*SetPasswordResp, error) {
	client := pb.NewAccountCenterClientClient(m.cli.Conn())
	return client.SetPassword(ctx, in, opts...)
}

func (m *defaultAccountCenterClient) CcnuLogin(ctx context.Context, in *CcnuLoginReq, opts ...grpc.CallOption) (*CcnuLoginResp, error) {
	client := pb.NewAccountCenterClientClient(m.cli.Conn())
	return client.CcnuLogin(ctx, in, opts...)
}

func (m *defaultAccountCenterClient) SetStudentID(ctx context.Context, in *SetStudentIDReq, opts ...grpc.CallOption) (*SetStudentIDResp, error) {
	client := pb.NewAccountCenterClientClient(m.cli.Conn())
	return client.SetStudentID(ctx, in, opts...)
}

func (m *defaultAccountCenterClient) SetEmail(ctx context.Context, in *SetEmailReq, opts ...grpc.CallOption) (*SetEmailResp, error) {
	client := pb.NewAccountCenterClientClient(m.cli.Conn())
	return client.SetEmail(ctx, in, opts...)
}