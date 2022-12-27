package store

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"gorm.io/gorm"
	model "postservice/postgres/gorm"
	pb "postservice/proto"
)

type PostStore interface {
	CreatePost(context.Context, *pb.Post, *pb.Post) error
	GetPost(context.Context, *pb.GetPostRequest, *pb.Post) error
	DeletePost(context.Context, *pb.DeletePostRequest, *empty.Empty) error
	UpdatePost(context.Context, *pb.UpdatePostRequest, *pb.Post) error
	ListPosts(context.Context, *empty.Empty, *pb.ListPostResponse) error
}

func NewPostsServer(
	db *gorm.DB,
) pb.PostsServiceHandler {

	//initial migration of databases: schema migration
	model.InitialMigrationUserProfile(db)
	return &PostPgStore{
		db: db,
	}
}
