package handler

import (
	_ "Blogs/docs"
	"Blogs/internal/service"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	admin := router.Group("/admin", h.middlewareAdmin)
	{
		admin.POST("/user", h.createUser)            //create
		admin.GET("/user", h.getAllUsers)            // getAll
		admin.POST("/user/:user_id", h.getUserById)  // getById
		admin.PUT("/user/:user_id", h.updateUser)    // update
		admin.DELETE("/user/:user_id", h.deleteUser) // delete
	}

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

		postComments := posts.Group("/:post_id/comments")
		{
			postComments.POST("/", h.createComment)
			postComments.GET("/", h.getAllComments)

		}
		comments := posts.Group("/comments")
		{
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
