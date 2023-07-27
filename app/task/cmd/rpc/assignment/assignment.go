package main

import (
	"flag"
	"fmt"

	"MuXiFresh-Be-2.0/app/task/cmd/rpc/assignment/internal/config"
	"MuXiFresh-Be-2.0/app/task/cmd/rpc/assignment/internal/server"
	"MuXiFresh-Be-2.0/app/task/cmd/rpc/assignment/internal/svc"
	"MuXiFresh-Be-2.0/app/task/cmd/rpc/assignment/pb"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/assignment.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterAssignmentClientServer(grpcServer, server.NewAssignmentClientServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
