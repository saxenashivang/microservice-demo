package store

import (
	"context"

	pb "github.com/saxenashivang/microservice-demo/src/userservice/proto"

	model "github.com/saxenashivang/microservice-demo/src/userservice/postgres/gorm"

	"github.com/golang/protobuf/ptypes/empty"
	"go-micro.dev/v4/logger"
	"gorm.io/gorm"
)

type UserPgStore struct {
	db *gorm.DB
}

func (e *UserPgStore) GetUser(ctx context.Context, req *pb.GetUserRequest, res *pb.User) error {
	logger.Infof("Received GetUser.Call request: %v", req)
	user := &pb.User{
		Id:            "1",
		Name:          "3",
		UserName:      "4",
		Email:         "5",
		PhoneNumber:   "6",
		ProfilePicUrl: "7",
		CreatedOn:     nil,
	}
	res.Id = user.Id
	res.Email = user.Email
	res.CreatedOn = user.CreatedOn
	res.Name = user.Name
	res.UserName = user.UserName
	res.ProfilePicUrl = user.ProfilePicUrl
	logger.Infof("Payload: %v", res)
	return nil
}
func (e *UserPgStore) CreateUser(ctx context.Context, req *pb.CreateUserRequest, res *pb.User) error {
	out := &model.User{
		ID:          1,
		Name:        req.User.Name,
		UserName:    req.User.UserName,
		Email:       req.User.Email,
		PhoneNumber: req.User.PhoneNumber,
	}
	e.db.Create(&out)
	return nil
}
func (e *UserPgStore) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest, res *empty.Empty) error {
	return nil
}
