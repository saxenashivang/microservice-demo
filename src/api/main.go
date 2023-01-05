package main

import (
	routers "api/router"

	"github.com/go-micro/plugins/v4/client/grpc"
	"go-micro.dev/v4"
)

func main() {
	srv := micro.NewService(
		micro.Client(grpc.NewClient()),
	)
	srv.Init()
	routers.Api(srv)
}
