package store

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"go-micro.dev/v4/logger"
	"gorm.io/gorm"
	pb "userservice/proto"
)

type UserPgStore struct {
	db *gorm.DB
}

func (e *UserPgStore) GetUser(ctx context.Context, req *pb.GetUserRequest, res *pb.User) error {
	logger.Infof("Received GetUser.Call request: %v", req)
	res = &pb.User{
		Id:            "1",
		Name:          "3",
		UserName:      "4",
		Email:         "5",
		PhoneNumber:   "6",
		ProfilePicUrl: "7",
		CreatedOn:     nil,
	}
	return nil
}
func (e *UserPgStore) CreateUser(ctx context.Context, req *pb.CreateUserRequest, res *pb.User) error {
	res = &pb.User{
		Id:            req.User.Id,
		Name:          req.User.Name,
		UserName:      req.User.UserName,
		Email:         req.User.Email,
		PhoneNumber:   req.User.PhoneNumber,
		ProfilePicUrl: req.User.ProfilePicUrl,
		CreatedOn:     nil,
	}
	e.db.Select("Name", "UserName", "CreatedAt", "Email", "PhoneNumber", "ProfilePicUrl").Create(&res)
	return nil
}
func (e *UserPgStore) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest, res *empty.Empty) error {
	return nil
}
