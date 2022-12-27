package store

import (
	"context"
	"database/sql"
	"github.com/golang/protobuf/ptypes/empty"
	"go-micro.dev/v4/logger"
	pb "userservice/proto"
)

type UserPgStore struct {
	*sql.DB
}

func (e *UserPgStore) GetUser(ctx context.Context, req *pb.GetUserRequest, res *pb.User) error {
	logger.Infof("Received Deployment.Call request: %v", req)
	res = &pb.User{
		Id:             "1",
		Sub:            "2",
		Name:           "3",
		UserName:       "4",
		Email:          "5",
		PhoneNumber:    "6",
		ProfilePicUrl:  "7",
		TokenValidTill: nil,
	}
	return nil
}
func (e *UserPgStore) CreateUser(ctx context.Context, req *pb.CreateUserRequest, res *pb.User) error {
	return nil
}
func (e *UserPgStore) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest, res *empty.Empty) error {
	return nil
}
