// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/postservice.proto

package postservice

import (
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	empty "github.com/golang/protobuf/ptypes/empty"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	_ "google.golang.org/genproto/protobuf/field_mask"
	proto "google.golang.org/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "go-micro.dev/v4/api"
	client "go-micro.dev/v4/client"
	server "go-micro.dev/v4/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for PostsService service

func NewPostsServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{
		{
			Name:    "PostsService.CreatePost",
			Path:    []string{"/v1/post"},
			Method:  []string{"POST"},
			Handler: "rpc",
		},
		{
			Name:    "PostsService.GetPost",
			Path:    []string{"/v1/post/{id}"},
			Method:  []string{"GET"},
			Handler: "rpc",
		},
		{
			Name:    "PostsService.DeletePost",
			Path:    []string{"/v1/post/{id}"},
			Method:  []string{"DELETE"},
			Handler: "rpc",
		},
		{
			Name:    "PostsService.UpdatePost",
			Path:    []string{"/v1/post/{id}"},
			Method:  []string{"PUT"},
			Handler: "rpc",
		},
		{
			Name:    "PostsService.ListPosts",
			Path:    []string{"/v1/post"},
			Method:  []string{"GET"},
			Handler: "rpc",
		},
	}
}

// Client API for PostsService service

type PostsService interface {
	CreatePost(ctx context.Context, in *Post, opts ...client.CallOption) (*Post, error)
	GetPost(ctx context.Context, in *GetPostRequest, opts ...client.CallOption) (*Post, error)
	DeletePost(ctx context.Context, in *DeletePostRequest, opts ...client.CallOption) (*empty.Empty, error)
	UpdatePost(ctx context.Context, in *UpdatePostRequest, opts ...client.CallOption) (*Post, error)
	ListPosts(ctx context.Context, in *empty.Empty, opts ...client.CallOption) (*ListPostResponse, error)
}

type postsService struct {
	c    client.Client
	name string
}

func NewPostsService(name string, c client.Client) PostsService {
	return &postsService{
		c:    c,
		name: name,
	}
}

func (c *postsService) CreatePost(ctx context.Context, in *Post, opts ...client.CallOption) (*Post, error) {
	req := c.c.NewRequest(c.name, "PostsService.CreatePost", in)
	out := new(Post)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postsService) GetPost(ctx context.Context, in *GetPostRequest, opts ...client.CallOption) (*Post, error) {
	req := c.c.NewRequest(c.name, "PostsService.GetPost", in)
	out := new(Post)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postsService) DeletePost(ctx context.Context, in *DeletePostRequest, opts ...client.CallOption) (*empty.Empty, error) {
	req := c.c.NewRequest(c.name, "PostsService.DeletePost", in)
	out := new(empty.Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postsService) UpdatePost(ctx context.Context, in *UpdatePostRequest, opts ...client.CallOption) (*Post, error) {
	req := c.c.NewRequest(c.name, "PostsService.UpdatePost", in)
	out := new(Post)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postsService) ListPosts(ctx context.Context, in *empty.Empty, opts ...client.CallOption) (*ListPostResponse, error) {
	req := c.c.NewRequest(c.name, "PostsService.ListPosts", in)
	out := new(ListPostResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for PostsService service

type PostsServiceHandler interface {
	CreatePost(context.Context, *Post, *Post) error
	GetPost(context.Context, *GetPostRequest, *Post) error
	DeletePost(context.Context, *DeletePostRequest, *empty.Empty) error
	UpdatePost(context.Context, *UpdatePostRequest, *Post) error
	ListPosts(context.Context, *empty.Empty, *ListPostResponse) error
}

func RegisterPostsServiceHandler(s server.Server, hdlr PostsServiceHandler, opts ...server.HandlerOption) error {
	type postsService interface {
		CreatePost(ctx context.Context, in *Post, out *Post) error
		GetPost(ctx context.Context, in *GetPostRequest, out *Post) error
		DeletePost(ctx context.Context, in *DeletePostRequest, out *empty.Empty) error
		UpdatePost(ctx context.Context, in *UpdatePostRequest, out *Post) error
		ListPosts(ctx context.Context, in *empty.Empty, out *ListPostResponse) error
	}
	type PostsService struct {
		postsService
	}
	h := &postsServiceHandler{hdlr}
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "PostsService.CreatePost",
		Path:    []string{"/v1/post"},
		Method:  []string{"POST"},
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "PostsService.GetPost",
		Path:    []string{"/v1/post/{id}"},
		Method:  []string{"GET"},
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "PostsService.DeletePost",
		Path:    []string{"/v1/post/{id}"},
		Method:  []string{"DELETE"},
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "PostsService.UpdatePost",
		Path:    []string{"/v1/post/{id}"},
		Method:  []string{"PUT"},
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "PostsService.ListPosts",
		Path:    []string{"/v1/post"},
		Method:  []string{"GET"},
		Handler: "rpc",
	}))
	return s.Handle(s.NewHandler(&PostsService{h}, opts...))
}

type postsServiceHandler struct {
	PostsServiceHandler
}

func (h *postsServiceHandler) CreatePost(ctx context.Context, in *Post, out *Post) error {
	return h.PostsServiceHandler.CreatePost(ctx, in, out)
}

func (h *postsServiceHandler) GetPost(ctx context.Context, in *GetPostRequest, out *Post) error {
	return h.PostsServiceHandler.GetPost(ctx, in, out)
}

func (h *postsServiceHandler) DeletePost(ctx context.Context, in *DeletePostRequest, out *empty.Empty) error {
	return h.PostsServiceHandler.DeletePost(ctx, in, out)
}

func (h *postsServiceHandler) UpdatePost(ctx context.Context, in *UpdatePostRequest, out *Post) error {
	return h.PostsServiceHandler.UpdatePost(ctx, in, out)
}

func (h *postsServiceHandler) ListPosts(ctx context.Context, in *empty.Empty, out *ListPostResponse) error {
	return h.PostsServiceHandler.ListPosts(ctx, in, out)
}