package main

import (
	"log"

	httpServer "github.com/go-micro/plugins/v4/server/http"
	"go-micro.dev/v4"

	"github.com/gin-gonic/gin"
	"go-micro.dev/v4/registry"
	"go-micro.dev/v4/server"
)

const (
	SERVER_NAME = "post-server-http" // server name
)

func main() {

	srv := httpServer.NewServer(
		server.Name(SERVER_NAME),
		server.Address(":8080"),
	)

	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(gin.Recovery())

	// register router
	demo := newPostService()
	demo.InitRouter(router)

	hd := srv.NewHandler(router)
	if err := srv.Handle(hd); err != nil {
		log.Fatalln(err)
	}

	service := micro.NewService(
		micro.Server(srv),
		micro.Registry(registry.NewRegistry()),
	)
	service.Init()
	service.Run()
}

// demo
type postService struct{}

func newPostService() *postService {
	return &postService{}
}

func (a *postService) InitRouter(router *gin.Engine) {
	router.GET("/post", a.GetPost)
}

func (a *postService) GetPost(c *gin.Context) {
	c.JSON(200, gin.H{"msg": "call go-micro v3 http server success"})
}
