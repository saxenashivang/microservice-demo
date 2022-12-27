package handler

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"go-micro.dev/v4/logger"
	"userservice/postgres/store"
	pb "userservice/proto"
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
	return nil
}
func (e *UserService) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest, res *empty.Empty) error {
	return nil
}
