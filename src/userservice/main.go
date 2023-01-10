package main

import (
	"context"
	"sync"

	"github.com/saxenashivang/microservice-demo/src/userservice/handler"
	pb "github.com/saxenashivang/microservice-demo/src/userservice/proto"

	ot "github.com/go-micro/plugins/v4/wrapper/trace/opentracing"
	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"
	"go-micro.dev/v4/server"

	"github.com/go-micro/cli/debug/trace/jaeger"

	"github.com/saxenashivang/microservice-demo/src/userservice/config"

	grpcc "github.com/go-micro/plugins/v4/client/grpc"
	grpcs "github.com/go-micro/plugins/v4/server/grpc"
)

var (
	service = "userservice"
	version = "latest"
)

func main() {
	// Load conigurations
	if err := config.Load(); err != nil {
		logger.Fatal(err)
	}
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
		micro.Address(":9001"),
	)
	srv.Server().Init(
		server.Wait(&wg),
	)

	ctx = server.NewContext(ctx, srv.Server())

	//check DB connection and migration
	// dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	// db, err := postgres.NewDB(dsn)
	if err != nil {
		logger.Infof("Unable to connect to DB %s", err)
	}
	// store := store.NewPostStore(db)
	// Register handler
	if err := pb.RegisterUserServiceHandler(srv.Server(), new(handler.UserService)); err != nil {
		logger.Fatal(err)
	}

	// if err := pb.RegisterUserServiceServer(grpcs.NewServer()); err != nil {

	// }

	if err := pb.RegisterHealthHandler(srv.Server(), new(handler.Health)); err != nil {
		logger.Fatal(err)
	}

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
