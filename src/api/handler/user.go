package handler

import (
	"github.com/gin-gonic/gin"
	"go-micro.dev/v4"
	userpb "userservice/proto"
)

type UserHandler struct{}

var user UserHandler

func (p *PostHandler) connUserServiceViaGRPC() userpb.UserService {
	srv := micro.NewService()

	srv.Init()

	// Create client
	c := userpb.NewUserService("userservice", srv.Client())
	return c
}
func GetUser(c *gin.Context) {
	conn := post.connUserServiceViaGRPC()
	res, err := conn.GetUser(c, &userpb.GetUserRequest{Id: "1"})
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"status": false, "error": err})
		return
	}
	c.JSON(200, res)
}

func CreateUser(c *gin.Context) {
	conn := post.connUserServiceViaGRPC()
	res, err := conn.CreateUser(c, &userpb.CreateUserRequest{User: &userpb.User{
		Id:            "1",
		Name:          "Shivang",
		UserName:      "shivang",
		Email:         "Sh@gmail.com",
		PhoneNumber:   "787822222",
		ProfilePicUrl: "",
		CreatedOn:     nil,
	}})
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"status": false, "error": err})
		return
	}
	c.JSON(200, res)
}
