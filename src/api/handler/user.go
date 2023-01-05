package handler

import (
	"context"
	"fmt"
	"net/http"
	userpb "userservice/proto"

	"github.com/gin-gonic/gin"
	"github.com/go-micro/plugins/v4/client/grpc"
	"go-micro.dev/v4"
)

type UserHandler struct{}

var user UserHandler

func (p *UserHandler) connUserServiceViaGRPC() userpb.UserService {
	srv := micro.NewService(
		micro.Client(grpc.NewClient()),
	)
	srv.Init()
	// c := srv.Client()

	// Create client
	c := userpb.NewUserService("userservice", srv.Client())
	return c
}

//	func (p *UserHandler) conngRPC() userpb.UserService {
//		conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
//		if err != nil {
//			log.Fatalf("did not connect: %v", err)
//		}
//		defer conn.Close()
//		client := userpb.NewUserService("userservice", nil)
//		return client
//	}
func GetUser(c *gin.Context) {
	conn := user.connUserServiceViaGRPC()
	res, err := conn.GetUser(context.Background(), &userpb.GetUserRequest{Id: "1"})
	fmt.Println(res, err)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"status": false, "error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"result": fmt.Sprint(res),
	})
}

func CreateUser(c *gin.Context) {
	conn := user.connUserServiceViaGRPC()
	res, err := conn.CreateUser(c, &userpb.CreateUserRequest{User: &userpb.User{
		Id:            "1",
		Name:          "Shivang",
		UserName:      "shivang",
		Email:         "Sh@gmail.com",
		PhoneNumber:   "787822222",
		ProfilePicUrl: "",
		CreatedOn:     nil,
	}})
	fmt.Println(res, err)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"status": false, "error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"result": fmt.Sprint(res),
	})
}
