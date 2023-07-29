// Code generated by goctl. DO NOT EDIT.
// Source: schedule.proto

package scheduleclient

import (
	"context"

	"MuXiFresh-Be-2.0/app/schedule/rpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	CheckReq   = pb.CheckReq
	CheckResp  = pb.CheckResp
	CreateReq  = pb.CreateReq
	CreateResp = pb.CreateResp

	ScheduleClient interface {
		Check(ctx context.Context, in *CheckReq, opts ...grpc.CallOption) (*CheckResp, error)
		Create(ctx context.Context, in *CreateReq, opts ...grpc.CallOption) (*CreateResp, error)
	}

	defaultScheduleClient struct {
		cli zrpc.Client
	}
)

func NewScheduleClient(cli zrpc.Client) ScheduleClient {
	return &defaultScheduleClient{
		cli: cli,
	}
}

func (m *defaultScheduleClient) Check(ctx context.Context, in *CheckReq, opts ...grpc.CallOption) (*CheckResp, error) {
	client := pb.NewScheduleClientClient(m.cli.Conn())
	return client.Check(ctx, in, opts...)
}

func (m *defaultScheduleClient) Create(ctx context.Context, in *CreateReq, opts ...grpc.CallOption) (*CreateResp, error) {
	client := pb.NewScheduleClientClient(m.cli.Conn())
	return client.Create(ctx, in, opts...)
}
