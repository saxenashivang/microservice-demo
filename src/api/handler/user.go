package handler

import (
	"context"
	"fmt"
	"net/http"
	userpb "userservice/proto"

	"github.com/gin-gonic/gin"
	"go-micro.dev/v4"
)

type UserHandler struct {
	con userpb.UserService
}

func NewUserHandler(srv micro.Service) *UserHandler {
	c := userpb.NewUserService("userservice", srv.Client())
	return &UserHandler{
		con: c,
	}
}

func (e *UserHandler) GetUser(c *gin.Context) {
	res, err := e.con.GetUser(context.Background(), &userpb.GetUserRequest{Id: "1"})
	fmt.Println(res, err)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"status": false, "error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"result": fmt.Sprint(res),
	})
}

func (e *UserHandler) CreateUser(c *gin.Context) {
	res, err := e.con.CreateUser(c, &userpb.CreateUserRequest{User: &userpb.User{
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
