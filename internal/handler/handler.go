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
		postsAuthor.GET("/my", h.getMyAllPosts)
		postsAuthor.POST("/", h.createPost)

		myPostsAuthor := postsAuthor.Group("/:post_id", h.checkOwnershipPost)
		{
			myPostsAuthor.PUT("/", h.updatePost)
			myPostsAuthor.DELETE("/", h.deletePost)
		}
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

			myComments := comments.Group("/:comment_id", h.checkOwnershipComment)
			{
				myComments.PUT("/", h.updateComment)
				myComments.DELETE("/", h.deleteComment)
			}
		}

		likes := posts.Group("/:post_id")
		{
			likes.POST("/like", h.likePost)
			likes.POST("/unlike", h.unlikePost)
		}
	}
	return router
}
