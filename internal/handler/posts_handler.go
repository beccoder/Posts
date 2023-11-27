package handler

import (
	"Blogs"
	"Blogs/internal/handler/http"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// createPost godoc
// @Summary 			Create a new Post
// @Security 			ApiKeyAuth
// @Tags 				Posts
// @Description 		Create a new Post with provided data
// @ID 					create-post
// @Accept 				json
// @Produce 			json
// @Param 				input body Blogs.CreatePostRequest true "post data"
// @Success 			200 {object} http.Response ""
// @Failure 			400,404 {object} http.Response
// @Failure				500 {object} http.Response
// @Failure 			default {object} http.Response
// @Router 				/posts [post]
func (h *Handler) createPost(c *gin.Context) {
	userId, err := h.getUserId(c)
	if err != nil {
		http.HandleResponse(c, http.NotFound, err.Error())
		return
	}

	var input Blogs.CreatePostRequest
	if err := c.BindJSON(&input); err != nil {
		http.HandleResponse(c, http.BadRequest, err.Error())
		return
	}

	createPostData := Blogs.PostModel{
		AuthorsId: userId,
		Title:     input.Title,
		Text:      input.Text,
	}

	id, err := h.services.CreatePosts(createPostData)
	if err != nil {
		http.HandleResponse(c, http.InternalServerError, err.Error())
		return
	}

	http.HandleResponse(c, http.Created, map[string]interface{}{
		"id": id,
	})
}

// getAllPosts godoc
// @Summary 			Get All Posts
// @Security 			ApiKeyAuth
// @Tags 				Posts
// @Description 		Get a list of all posts
// @ID 					get-posts
// @Accept 				json
// @Produce 			json
// @Success 			200 {object} http.Response
// @Failure 			400,404 {object} http.Response
// @Failure				500 {object} http.Response
// @Failure 			default {object} http.Response
// @Router 				/posts [get]
func (h *Handler) getAllPosts(c *gin.Context) {
	posts, err := h.services.GetAllPosts()
	if err != nil {
		http.HandleResponse(c, http.InternalServerError, err.Error())
		return
	}

	http.HandleResponse(c, http.OK, posts)
}

// getMyAllPosts godoc
// @Summary 			Get All My Posts
// @Security 			ApiKeyAuth
// @Tags 				Posts
// @Description 		Get a list of all my posts
// @ID 					get-my-posts
// @Accept 				json
// @Produce 			json
// @Success 			200 {object} http.Response
// @Failure 			400,404 {object} http.Response
// @Failure				500 {object} http.Response
// @Failure 			default {object} http.Response
// @Router 				/posts/my [get]
func (h *Handler) getMyAllPosts(c *gin.Context) {
	userId, err := h.getUserId(c)
	if err != nil {
		http.HandleResponse(c, http.NotFound, err.Error())
		return
	}
	posts, err := h.services.GetMyAllPosts(userId)
	if err != nil {
		http.HandleResponse(c, http.InternalServerError, err.Error())
		return
	}

	http.HandleResponse(c, http.OK, posts)
}

// getPostById godoc
// @Summary 			Get a Post by ID
// @Security 			ApiKeyAuth
// @Tags 				Posts
// @Description 		Get a single Post by providing its ID
// @ID 					get-post
// @Accept 				json
// @Produce 			json
// @Param 				post_id path string true "Post ID"
// @Success 			200 {object} http.Response
// @Failure 			400,404 {object} http.Response
// @Failure				500 {object} http.Response
// @Failure 			default {object} http.Response
// @Router 				/posts/{post_id} [get]
func (h *Handler) getPostById(c *gin.Context) {
	postId, err := primitive.ObjectIDFromHex(c.Param("post_id"))
	if err != nil {
		http.HandleResponse(c, http.InvalidArgument, err.Error())
		return
	}

	post, err := h.services.GetPostById(postId)
	if err != nil {
		if err.Error() == "no posts exist" {
			http.HandleResponse(c, http.NotFound, err.Error())
			return
		}
		http.HandleResponse(c, http.InternalServerError, err.Error())
		return
	}

	http.HandleResponse(c, http.OK, post)
}

// updatePost godoc
// @Summary 			Update an existing Post data
// @Security 			ApiKeyAuth
// @Tags 				Posts
// @Description 		Update an existing Post data with new data
// @ID 					update-post
// @Accept 				json
// @Produce 			json
// @Param 				post_id path string true "Post ID"
// @Param 				new_input body Blogs.UpdatePostRequest true "new post info"
// @Success 			200 {object} http.Response
// @Failure 			400,404 {object} http.Response
// @Failure				500 {object} http.Response
// @Failure 			default {object} http.Response
// @Router 				/posts/{post_id} [put]
func (h *Handler) updatePost(c *gin.Context) {
	postId, err := primitive.ObjectIDFromHex(c.Param("post_id"))
	if err != nil {
		http.HandleResponse(c, http.InvalidArgument, err.Error())
		return
	}

	var input Blogs.UpdatePostRequest
	if err := c.BindJSON(&input); err != nil {
		http.HandleResponse(c, http.BadRequest, err.Error())
		return
	}

	err = h.services.UpdatePost(postId, input)
	if err != nil {
		http.HandleResponse(c, http.InternalServerError, err.Error())
		return
	}

	http.HandleResponse(c, http.OK, statusResponse{"ok"})
}

// deletePost godoc
// @Summary				Delete an existing Post
// @Security 			ApiKeyAuth
// @Tags 				Posts
// @Description			Delete an existing Post by ID
// @ID 					delete-post
// @Accept 				json
// @Produce 			json
// @Param 				post_id path string true "Post ID"
// @Success 			200 {object} http.Response
// @Failure 			400,404 {object} http.Response
// @Failure				500 {object} http.Response
// @Failure 			default {object} http.Response
// @Router 				/posts/{post_id} [delete]
func (h *Handler) deletePost(c *gin.Context) {
	postId, err := primitive.ObjectIDFromHex(c.Param("post_id"))
	if err != nil {
		http.HandleResponse(c, http.BadRequest, err.Error())
		return
	}

	err = h.services.DeletePost(postId)
	if err != nil {
		http.HandleResponse(c, http.InternalServerError, err.Error())
		return
	}

	http.HandleResponse(c, http.OK, statusResponse{"ok"})
}

// createComment godoc
// @Summary 			Create a new Comment
// @Security 			ApiKeyAuth
// @Tags 				Posts
// @Description 		Create a new Comment with provided data
// @ID 					comment-on-post
// @Accept 				json
// @Produce 			json
// @Param 				post_id path string true "Post ID"
// @Param 				input body Blogs.CreateCommentRequest true "comment data"
// @Success 			200 {object} http.Response ""
// @Failure 			400,404 {object} http.Response
// @Failure				500 {object} http.Response
// @Failure 			default {object} http.Response
// @Router 				/posts/{post_id}/comments [post]
func (h *Handler) createComment(c *gin.Context) {
	userId, err := h.getUserId(c)
	if err != nil {
		http.HandleResponse(c, http.InvalidArgument, err.Error())
		return
	}
	postId, err := primitive.ObjectIDFromHex(c.Param("post_id"))
	if err != nil {
		http.HandleResponse(c, http.InvalidArgument, err.Error())
		return
	}
	var input Blogs.CreateCommentRequest
	if err := c.BindJSON(&input); err != nil {
		http.HandleResponse(c, http.BadRequest, err.Error())
		return
	}

	createCommentData := Blogs.CommentModel{
		PostId:         postId,
		CommentedById:  userId,
		ReplyCommentId: input.ReplyCommentId,
		Comment:        input.Comment,
	}

	id, err := h.services.CreateComment(createCommentData)
	if err != nil {
		http.HandleResponse(c, http.InternalServerError, err.Error())
		return
	}

	http.HandleResponse(c, http.Created, map[string]interface{}{
		"id": id,
	})
}

// getAllComments godoc
// @Summary 			Get All Comments
// @Security 			ApiKeyAuth
// @Tags 				Posts
// @Description 		Get a list of all comments related to this post
// @ID 					get-post-comments
// @Accept 				json
// @Produce 			json
// @Param 				post_id path string true "Post ID"
// @Success 			200 {object} http.Response
// @Failure 			400,404 {object} http.Response
// @Failure				500 {object} http.Response
// @Failure 			default {object} http.Response
// @Router 				/posts/{post_id}/comments [get]
func (h *Handler) getAllComments(c *gin.Context) {
	postId, err := primitive.ObjectIDFromHex(c.Param("post_id"))
	if err != nil {
		http.HandleResponse(c, http.InvalidArgument, err.Error())
		return
	}
	comments, err := h.services.GetAllComments(postId)
	if err != nil {
		http.HandleResponse(c, http.InternalServerError, err.Error())
		return
	}

	http.HandleResponse(c, http.OK, comments)
}

// getCommentById godoc
// @Summary 			Get a Comment by ID
// @Security 			ApiKeyAuth
// @Tags 				Posts
// @Description 		Get a single Comment of the post by providing its ID
// @ID 					get-post-comment
// @Accept 				json
// @Produce 			json
// @Param 				comment_id path string true "Comment ID"
// @Success 			200 {object} http.Response
// @Failure 			400,404 {object} http.Response
// @Failure				500 {object} http.Response
// @Failure 			default {object} http.Response
// @Router 				/posts/comments/{comment_id} [get]
func (h *Handler) getCommentById(c *gin.Context) {
	commentId, err := primitive.ObjectIDFromHex(c.Param("comment_id"))
	if err != nil {
		http.HandleResponse(c, http.InvalidArgument, err.Error())
		return
	}

	comment, err := h.services.GetCommentById(commentId)
	if err != nil {
		if err.Error() == "no comments exist" {
			http.HandleResponse(c, http.NotFound, err.Error())
			return
		}
		http.HandleResponse(c, http.InternalServerError, err.Error())
		return
	}

	http.HandleResponse(c, http.OK, comment)
}

// updateComment godoc
// @Summary 			Update an existing Comment data
// @Security 			ApiKeyAuth
// @Tags 				Posts
// @Description 		Update an existing Comment data of the post with new data
// @ID 					update-post-comment
// @Accept 				json
// @Produce 			json
// @Param 				comment_id path string true "Comment ID"
// @Param 				new_input body Blogs.UpdateCommentRequest true "new comment data"
// @Success 			200 {object} http.Response
// @Failure 			400,404 {object} http.Response
// @Failure				500 {object} http.Response
// @Failure 			default {object} http.Response
// @Router 				/posts/comments/{comment_id} [put]
func (h *Handler) updateComment(c *gin.Context) {
	commentId, err := primitive.ObjectIDFromHex(c.Param("comment_id"))
	if err != nil {
		http.HandleResponse(c, http.InvalidArgument, err.Error())
		return
	}

	var input Blogs.UpdateCommentRequest
	if err := c.BindJSON(&input); err != nil {
		http.HandleResponse(c, http.BadRequest, err.Error())
		return
	}

	err = h.services.UpdateComment(commentId, input)
	if err != nil {
		http.HandleResponse(c, http.InternalServerError, err.Error())
		return
	}

	http.HandleResponse(c, http.OK, statusResponse{"ok"})
}

// deleteComment godoc
// @Summary				Delete an existing Comment of the post
// @Security 			ApiKeyAuth
// @Tags 				Posts
// @Description			Delete an existing Comment of the post by ID
// @ID 					delete-post-comment
// @Accept 				json
// @Produce 			json
// @Param 				comment_id path string true "Comment ID"
// @Success 			200 {object} http.Response
// @Failure 			400,404 {object} http.Response
// @Failure				500 {object} http.Response
// @Failure 			default {object} http.Response
// @Router 				/posts/comments/{comment_id} [delete]
func (h *Handler) deleteComment(c *gin.Context) {
	commentId, err := primitive.ObjectIDFromHex(c.Param("comment_id"))
	if err != nil {
		http.HandleResponse(c, http.InvalidArgument, err.Error())
		return
	}

	err = h.services.DeleteComment(commentId)
	if err != nil {
		http.HandleResponse(c, http.InternalServerError, err.Error())
		return
	}

	http.HandleResponse(c, http.OK, statusResponse{"ok"})
}

// likePost godoc
// @Summary 			Like the post
// @Security 			ApiKeyAuth
// @Tags 				Posts
// @Description 		Like the post
// @ID 					like-post
// @Accept 				json
// @Produce 			json
// @Param 				post_id path string true "Post ID"
// @Success 			200 {object} http.Response
// @Failure 			400,404 {object} http.Response
// @Failure				500 {object} http.Response
// @Failure 			default {object} http.Response
// @Router 				/posts/{post_id}/like [post]
func (h *Handler) likePost(c *gin.Context) {
	userId, err := h.getUserId(c)
	if err != nil {
		http.HandleResponse(c, http.InvalidArgument, err.Error())
		return
	}
	postId, err := primitive.ObjectIDFromHex(c.Param("post_id"))
	if err != nil {
		http.HandleResponse(c, http.InvalidArgument, err.Error())
		return
	}

	err = h.services.AddLike(postId, userId)
	if err != nil {
		if err.Error() == "already liked" {
			http.HandleResponse(c, http.RequestConflict, err.Error())
			return
		}
		if err.Error() == "no posts exist" {
			http.HandleResponse(c, http.NotFound, err.Error())
			return
		}
		http.HandleResponse(c, http.InternalServerError, err.Error())
		return
	}

	http.HandleResponse(c, http.OK, statusResponse{"ok"})
}

// unlikePost godoc
// @Summary 			Unlike the post
// @Security 			ApiKeyAuth
// @Tags 				Posts
// @Description 		Unlike the post
// @ID 					unlike-post
// @Accept 				json
// @Produce 			json
// @Param 				post_id path string true "Post ID"
// @Success 			200 {object} http.Response
// @Failure 			400,404 {object} http.Response
// @Failure				500 {object} http.Response
// @Failure 			default {object} http.Response
// @Router 				/posts/{post_id}/unlike [post]
func (h *Handler) unlikePost(c *gin.Context) {
	userId, err := h.getUserId(c)
	if err != nil {
		http.HandleResponse(c, http.InvalidArgument, err.Error())
		return
	}
	postId, err := primitive.ObjectIDFromHex(c.Param("post_id"))
	if err != nil {
		http.HandleResponse(c, http.InvalidArgument, err.Error())
		return
	}

	err = h.services.UnlikePost(postId, userId)
	if err != nil {
		http.HandleResponse(c, http.InternalServerError, err.Error())
		return
	}

	http.HandleResponse(c, http.OK, statusResponse{"ok"})
}
