package routers

import (
	"api/handler"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Api() *gin.Engine {
	router := gin.Default()
	/* Public API */
	router.GET("/api/v1/ping", handler.Pong)
	//router.GET("/api/v1/get-post", handler.GetPost)
	router.POST("/api/v1/create-user", handler.CreateUser)
	router.GET("/api/v1/get-user", handler.GetUser)
	//router.GET("/api/v1/delete-user", handler.GetPost)

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

	router.Run(fmt.Sprintf("0.0.0.0:%s", "8080"))

	return router
}
