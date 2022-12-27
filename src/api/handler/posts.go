package handler

import (
	"github.com/gin-gonic/gin"
	"go-micro.dev/v4"
	pb "postservice/proto"
)

type PostHandler struct{}

var post PostHandler

func (p *PostHandler) connGrpc() pb.PostsService {
	srv := micro.NewService()

	srv.Init()

	// Create client
	c := pb.NewPostsService("post-service", srv.Client())
	return c
}

func GetPost(c *gin.Context) {
	conn := post.connGrpc()
	res, err := conn.GetPost(c, &pb.GetPostRequest{Id: "1"})
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"status": false, "error": err})
		return
	}
	c.JSON(200, res)
}
