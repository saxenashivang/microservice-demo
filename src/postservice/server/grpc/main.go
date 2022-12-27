package main

import (
	"context"
	"log"

	"go-micro.dev/v4"
	"go-micro.dev/v4/client"
	"go-micro.dev/v4/registry"
	pb "post-service/proto"
)

func main() {
	CallGrpcByHttp()
}

func CallGrpcByHttp() {
	// create a new service
	service := micro.NewService(
		micro.Registry(registry.NewRegistry()),
	)
	// parse command line flags
	service.Init()
	c := service.Client()

	req := &pb.GetPostRequest{Id: "1"}
	// create request/response
	request := client.NewRequest("go.micro.srv.post-service", "PostsService.GetPost", req)

	response := new(pb.Post)
	// call service
	err := c.Call(context.Background(), request, response)
	log.Printf("err:%v response:%#v\n", err, response)
}
