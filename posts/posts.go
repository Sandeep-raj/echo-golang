package posts

import "github.com/labstack/echo/v4"

type Post struct {
	UserId int    `json:"userId"`
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func PostGroupHandler(post *echo.Group) {
	post.GET("/", GetPosts)
	post.GET("/userid/:userid", GetPostsByUserId)
	post.GET("/id/:id", GetPostsById)
}
