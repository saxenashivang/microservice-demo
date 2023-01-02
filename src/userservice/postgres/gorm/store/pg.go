package store

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"go-micro.dev/v4/logger"
	"gorm.io/gorm"
	pb "postservice/proto"
)

type PostPgStore struct {
	db *gorm.DB
}

func (e *PostPgStore) CreatePost(ctx context.Context, req *pb.Post, res *pb.Post) error {
	logger.Infof("Received PostService.CreatePost request: %v", req)
	return nil
}
func (e *PostPgStore) GetPost(ctx context.Context, req *pb.GetPostRequest, res *pb.Post) error {
	logger.Infof("Received PostService.GetPost request: %v", req)
	res.Id = "1"
	res.Title = "hooked"
	res.Description = "build habbit forming products"
	return nil
}
func (e *PostPgStore) DeletePost(ctx context.Context, req *pb.DeletePostRequest, res *empty.Empty) error {
	logger.Infof("Received PostService.CreatePost request: %v", req)
	return nil
}
func (e *PostPgStore) UpdatePost(ctx context.Context, req *pb.UpdatePostRequest, res *pb.Post) error {
	logger.Infof("Received PostService.CreatePost request: %v", req)
	return nil
}
func (e *PostPgStore) ListPosts(ctx context.Context, req *empty.Empty, res *pb.ListPostResponse) error {
	logger.Infof("Received PostService.CreatePost request: %v", req)
	return nil
}
