package handler

import (
	"Blogs/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		author := auth.Group("/author")
		{
			author.POST("/sign-in", h.signInAuthor)
			author.POST("/sign-up", h.signUpAuthor)
		}
		user := auth.Group("/user")
		{
			user.POST("/sign-in", h.signInUser)
			user.POST("/sign-up", h.signUpUser)
		}

	}

	postsAuthor := router.Group("/posts", h.middlewareAuthor)
	{
		postsAuthor.POST("/", h.createPost)
		postsAuthor.PUT("/:post_id", h.updatePost)
		postsAuthor.DELETE("/:post_id", h.deletePost)
	}

	posts := router.Group("/posts", h.middleware)
	{
		posts.GET("/", h.getAllPosts)
		posts.GET("/:post_id", h.getPostById)

		comments := posts.Group("/:post_id/comments")
		{
			comments.POST("/", h.createComment)
			comments.GET("/", h.getAllComments)
			comments.GET("/:comment_id", h.getCommentById)
			comments.PUT("/:comment_id", h.updateComment)
			comments.DELETE("/:comment_id", h.deleteComment)
		}

		likes := posts.Group("/:post_id/likes") // tricky case
		{
			likes.POST("/", h.createLike)
			likes.GET("/", h.getAllLikes)
			likes.GET("/:like_id", h.getLikeById)
			likes.DELETE("/:like_id", h.deleteLike)
		}
	}
	return router
}
