package main

import (
	"context"
	"sync"

	"postservice/handler"
	pb "postservice/proto"

	ot "github.com/go-micro/plugins/v4/wrapper/trace/opentracing"
	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"
	"go-micro.dev/v4/server"

	"github.com/go-micro/cli/debug/trace/jaeger"
	grpcc "github.com/go-micro/plugins/v4/client/grpc"
	grpcs "github.com/go-micro/plugins/v4/server/grpc"
)

var (
	service = "post-service"
	version = "latest"
)

func main() {
	// Create tracer
	tracer, closer, err := jaeger.NewTracer(
		jaeger.Name(service),
		jaeger.FromEnv(true),
		jaeger.GlobalTracer(true),
	)
	if err != nil {
		logger.Fatal(err)
	}
	defer closer.Close()

	wg := sync.WaitGroup{}
	ctx, cancel := context.WithCancel(context.Background())

	// Create service
	srv := micro.NewService(
		micro.Server(grpcs.NewServer()),
		micro.Client(grpcc.NewClient()),
		micro.BeforeStart(func() error {
			logger.Infof("Starting service %s", service)
			return nil
		}),
		micro.BeforeStop(func() error {
			logger.Infof("Shutting down service %s", service)
			cancel()
			return nil
		}),
		micro.AfterStop(func() error {
			wg.Wait()
			return nil
		}),
		micro.WrapCall(ot.NewCallWrapper(tracer)),
		micro.WrapClient(ot.NewClientWrapper(tracer)),
		micro.WrapHandler(ot.NewHandlerWrapper(tracer)),
		micro.WrapSubscriber(ot.NewSubscriberWrapper(tracer)),
	)
	srv.Init(
		micro.Name(service),
		micro.Version(version),
		micro.Address(":8082"),
	)
	srv.Server().Init(
		server.Wait(&wg),
	)

	ctx = server.NewContext(ctx, srv.Server())

	// Register handler
	if err := pb.RegisterPostsServiceHandler(srv.Server(), new(handler.PostService)); err != nil {
		logger.Fatal(err)
	}
	if err := pb.RegisterHealthHandler(srv.Server(), new(handler.Health)); err != nil {
		logger.Fatal(err)
	}

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
