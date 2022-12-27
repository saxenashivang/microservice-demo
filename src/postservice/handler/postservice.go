package handler

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"go-micro.dev/v4/logger"
	"postservice/postgres/gorm/store"

	pb "postservice/proto"
)

type PostService struct {
	Store store.PostStore
}

func (e *PostService) CreatePost(ctx context.Context, req *pb.Post, res *pb.Post) error {
	logger.Infof("Received PostService.CreatePost request: %v", req)
	err := req.Validate()
	if err != nil {
		return err
	}
	return e.Store.CreatePost(ctx, req, res)
}
func (e *PostService) GetPost(ctx context.Context, req *pb.GetPostRequest, res *pb.Post) error {
	logger.Infof("Received PostService.GetPost request: %v", req)
	res.Id = "1"
	res.Title = "hooked"
	res.Description = "build habbit forming products"
	return nil
}
func (e *PostService) DeletePost(ctx context.Context, req *pb.DeletePostRequest, res *empty.Empty) error {
	logger.Infof("Received PostService.CreatePost request: %v", req)
	return nil
}
func (e *PostService) UpdatePost(ctx context.Context, req *pb.UpdatePostRequest, res *pb.Post) error {
	logger.Infof("Received PostService.CreatePost request: %v", req)
	return nil
}
func (e *PostService) ListPosts(ctx context.Context, req *empty.Empty, res *pb.ListPostResponse) error {
	logger.Infof("Received PostService.CreatePost request: %v", req)
	return nil
}
