package main

import (
	"context"
	"time"

	pb "postservice/proto"

	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"
)

var (
	service = "post-service"
	version = "latest"
)

func main() {
	// Create service

	srv := micro.NewService()

	srv.Init()

	// Create client
	c := pb.NewPostsService(service, srv.Client())

	for {
		// Call service
		rsp, err := c.GetPost(context.Background(), &pb.GetPostRequest{Id: "John"})
		if err != nil {
			logger.Fatal(err)
		}

		logger.Info(rsp)

		time.Sleep(1 * time.Second)
	}
}
