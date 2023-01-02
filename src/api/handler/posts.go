package handler

type PostHandler struct{}

var post PostHandler

//func (p *PostHandler) connPostServiceViaGRPC() postpb.PostsService {
//	srv := micro.NewService()
//
//	srv.Init()
//
//	// Create client
//	c := postpb.NewPostsService("postservice", srv.Client())
//	return c
//}
//
//func GetPost(c *gin.Context) {
//	conn := post.connPostServiceViaGRPC()
//	res, err := conn.GetPost(c, &postpb.GetPostRequest{Id: "1"})
//	if err != nil {
//		c.AbortWithStatusJSON(400, gin.H{"status": false, "error": err})
//		return
//	}
//	c.JSON(200, res)
//}
