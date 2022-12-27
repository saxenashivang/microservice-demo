package store

import (
	"context"
	"database/sql"
	"github.com/golang/protobuf/ptypes/empty"
	pb "userservice/proto"
)

type UserStore interface {
	GetUser(ctx context.Context, in *pb.GetUserRequest, out *pb.User) error
	CreateUser(ctx context.Context, in *pb.CreateUserRequest, out *pb.User) error
	DeleteUser(ctx context.Context, in *pb.DeleteUserRequest, out *empty.Empty) error
}

func NewPostStore(db *sql.DB) UserStore {
	return &UserPgStore{db}
}
