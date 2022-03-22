package comments

import "github.com/labstack/echo/v4"

type Comment struct {
	PostId int    `json:"postId"`
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Body   string `json:"body"`
}

func CommentGroupHandler(comment *echo.Group) {
	comment.GET("/", GetComments)
	comment.GET("/postid/:postid", GetCommentsByPostId)
	comment.GET("/id/:id", GetCommentsById)
}
