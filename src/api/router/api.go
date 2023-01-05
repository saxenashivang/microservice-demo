package routers

import (
	"api/handler"
	"fmt"

	"github.com/gin-gonic/gin"
	"go-micro.dev/v4"
)

func Api(srv micro.Service) *gin.Engine {
	router := gin.Default()
	user := handler.NewUserHandler(srv)
	post := handler.NewPostHandler(srv)
	/* Public API */
	router.GET("/api/v1/ping", handler.Pong)
	router.POST("/api/v1/create-user", user.CreateUser)
	router.GET("/api/v1/get-user", user.GetUser)

	router.GET("/api/v1/get-post", post.GetPost)

	/* Private API */
	// middleware implementation gin based
	//router.Use(middleware.Authorization())
	//userR := router.Group("/api/v1/users")
	//{
	//	userR.GET("", handler.FindAllUser)
	//	userR.GET("/:id", handler.FineUserById)
	//	userR.PATCH("/:id", handler.UpdateUser)
	//}
	//
	//postR := router.Group("/api/v1/posts")
	//{
	//	postR.GET("", handler.ListPost)
	//	postR.POST("", handler.CreatePost)
	//}

	router.Run(fmt.Sprintf("0.0.0.0:%s", "8888"))

	return router
}
