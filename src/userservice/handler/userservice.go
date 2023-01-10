package handler

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/saxenashivang/microservice-demo/src/userservice/postgres/gorm/store"
	pb "github.com/saxenashivang/microservice-demo/src/userservice/proto"
	"go-micro.dev/v4/logger"
)

type UserService struct {
	Store store.UserPgStore
}

func (e *UserService) GetUser(ctx context.Context, req *pb.GetUserRequest, res *pb.User) error {
	logger.Infof("Calling Get User: %v", req)
	err := req.Validate()
	if err != nil {
		return err
	}
	return e.Store.GetUser(ctx, req, res)
}
func (e *UserService) CreateUser(ctx context.Context, req *pb.CreateUserRequest, res *pb.User) error {
	logger.Infof("Calling Create User: %v", req)
	err := req.Validate()
	if err != nil {
		return err
	}
	return e.Store.CreateUser(ctx, req, res)
}
func (e *UserService) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest, res *empty.Empty) error {
	logger.Infof("Calling Delete User: %v", req)
	err := req.Validate()
	if err != nil {
		return err
	}
	return e.Store.DeleteUser(ctx, req, res)
}
