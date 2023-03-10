package store

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	model "github.com/saxenashivang/microservice-demo/src/userservice/postgres/gorm"
	pb "github.com/saxenashivang/microservice-demo/src/userservice/proto"
	"gorm.io/gorm"
)

type UserStore interface {
	GetUser(ctx context.Context, in *pb.GetUserRequest, out *pb.User) error
	CreateUser(ctx context.Context, in *pb.CreateUserRequest, out *pb.User) error
	DeleteUser(ctx context.Context, in *pb.DeleteUserRequest, out *empty.Empty) error
}

func NewPostStore(db *gorm.DB) pb.UserServiceHandler {
	//initial migration of databases: schema migration
	model.InitialMigrationUserProfile(db)
	return &UserPgStore{
		db: db,
	}
}
