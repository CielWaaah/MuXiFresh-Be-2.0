// Code generated by goctl. DO NOT EDIT.
// Source: test.proto

package testclient

import (
	"context"

	"MuxiFresh2.0/MuXiFresh-Be-2.0/app/test/rpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	TestInfoReq  = pb.TestInfoReq
	TestInfoResp = pb.TestInfoResp
	TestReq      = pb.TestReq
	TestResp     = pb.TestResp
	UserTypeReq  = pb.UserTypeReq
	UserTypeResp = pb.UserTypeResp

	TestClient interface {
		UserTest(ctx context.Context, in *TestReq, opts ...grpc.CallOption) (*TestResp, error)
		CheckTestResult(ctx context.Context, in *TestInfoReq, opts ...grpc.CallOption) (*TestInfoResp, error)
		JudgeUserType(ctx context.Context, in *UserTypeReq, opts ...grpc.CallOption) (*UserTypeResp, error)
	}

	defaultTestClient struct {
		cli zrpc.Client
	}
)

func NewTestClient(cli zrpc.Client) TestClient {
	return &defaultTestClient{
		cli: cli,
	}
}

func (m *defaultTestClient) UserTest(ctx context.Context, in *TestReq, opts ...grpc.CallOption) (*TestResp, error) {
	client := pb.NewTestClientClient(m.cli.Conn())
	return client.UserTest(ctx, in, opts...)
}

func (m *defaultTestClient) CheckTestResult(ctx context.Context, in *TestInfoReq, opts ...grpc.CallOption) (*TestInfoResp, error) {
	client := pb.NewTestClientClient(m.cli.Conn())
	return client.CheckTestResult(ctx, in, opts...)
}

func (m *defaultTestClient) JudgeUserType(ctx context.Context, in *UserTypeReq, opts ...grpc.CallOption) (*UserTypeResp, error) {
	client := pb.NewTestClientClient(m.cli.Conn())
	return client.JudgeUserType(ctx, in, opts...)
}
