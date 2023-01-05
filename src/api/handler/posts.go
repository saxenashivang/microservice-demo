package handler

import (
	postpb "postservice/proto"

	"github.com/gin-gonic/gin"
	"go-micro.dev/v4"
)

type PostHandler struct {
	con postpb.PostsService
}

func NewPostHandler(srv micro.Service) *PostHandler {
	c := postpb.NewPostsService("post-service", srv.Client())
	return &PostHandler{
		con: c,
	}
}

func (e *PostHandler) GetPost(c *gin.Context) {
	res, err := e.con.GetPost(c, &postpb.GetPostRequest{Id: "1"})
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"status": false, "error": err})
		return
	}
	c.JSON(200, res)
}
