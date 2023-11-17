package handler

import (
	"Blogs"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
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
// @Success 			200 {integer} integer 1
// @Failure 			400,404 {object} errorResponse
// @Failure				500 {object} errorResponse
// @Failure 			default {object} errorResponse
// @Router 				/posts [post]
func (h *Handler) createPost(c *gin.Context) {
	userId, err := h.getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "id not found", err)
		return
	}

	var input Blogs.CreatePostRequest
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid type of input", err)
		return
	}

	createPostData := Blogs.PostModel{
		AuthorsId: userId,
		Title:     input.Title,
		Text:      input.Text,
	}

	id, err := h.services.CreatePosts(createPostData)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "Error while creating post", err)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type getAllPostsResponse struct {
	Data []Blogs.PostResponse `json:"data"`
}

// getAllPosts godoc
// @Summary 			Get All Posts
// @Security 			ApiKeyAuth
// @Tags 				Posts
// @Description 		Get a list of all posts
// @ID 					get-posts
// @Accept 				json
// @Produce 			json
// @Success 			200 {object} getAllPostsResponse
// @Failure 			400,404 {object} errorResponse
// @Failure				500 {object} errorResponse
// @Failure 			default {object} errorResponse
// @Router 				/posts [get]
func (h *Handler) getAllPosts(c *gin.Context) {
	posts, err := h.services.GetAllPosts()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "error while getting posts", err)
		return
	}

	c.JSON(http.StatusOK, getAllPostsResponse{
		Data: posts,
	})
}

// getMyAllPosts godoc
// @Summary 			Get All My Posts
// @Security 			ApiKeyAuth
// @Tags 				Posts
// @Description 		Get a list of all my posts
// @ID 					get-my-posts
// @Accept 				json
// @Produce 			json
// @Success 			200 {object} getAllPostsResponse
// @Failure 			400,404 {object} errorResponse
// @Failure				500 {object} errorResponse
// @Failure 			default {object} errorResponse
// @Router 				/posts/my [get]
func (h *Handler) getMyAllPosts(c *gin.Context) {
	userId, err := h.getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "id not found", err)
		return
	}
	posts, err := h.services.GetMyAllPosts(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "error while getting posts", err)
		return
	}

	c.JSON(http.StatusOK, getAllPostsResponse{
		Data: posts,
	})
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
// @Success 			200 {object} Blogs.PostResponse
// @Failure 			400,404 {object} errorResponse
// @Failure				500 {object} errorResponse
// @Failure 			default {object} errorResponse
// @Router 				/posts/{post_id} [get]
func (h *Handler) getPostById(c *gin.Context) {
	postId, err := primitive.ObjectIDFromHex(c.Param("post_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid id param", err)
		return
	}

	post, err := h.services.GetPostById(postId)
	if err != nil {
		if err.Error() == "no posts exist" {
			newErrorResponse(c, http.StatusInternalServerError, "No posts found", err)
			return
		}
		newErrorResponse(c, http.StatusInternalServerError, "Error in getting post", err)
		return
	}

	c.JSON(http.StatusOK, post)
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
// @Success 			200 {object} statusResponse
// @Failure 			400,404 {object} errorResponse
// @Failure				500 {object} errorResponse
// @Failure 			default {object} errorResponse
// @Router 				/posts/{post_id} [put]
func (h *Handler) updatePost(c *gin.Context) {
	postId, err := primitive.ObjectIDFromHex(c.Param("post_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid id param", err)
		return
	}

	var input Blogs.UpdatePostRequest
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid type of input", err)
		return
	}

	err = h.services.UpdatePost(postId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "Error while updating post", err)
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
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
// @Success 			200 {object} statusResponse
// @Failure 			400,404 {object} errorResponse
// @Failure				500 {object} errorResponse
// @Failure 			default {object} errorResponse
// @Router 				/posts/{post_id} [delete]
func (h *Handler) deletePost(c *gin.Context) {
	postId, err := primitive.ObjectIDFromHex(c.Param("post_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid id param", err)
		return
	}

	err = h.services.DeletePost(postId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "Error while deleting post", err)
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
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
// @Success 			200 {integer} integer 1
// @Failure 			400,404 {object} errorResponse
// @Failure				500 {object} errorResponse
// @Failure 			default {object} errorResponse
// @Router 				/posts/{post_id}/comments [post]
func (h *Handler) createComment(c *gin.Context) {
	userId, err := h.getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "id not found", err)
		return
	}
	postId, err := primitive.ObjectIDFromHex(c.Param("post_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid id param", err)
		return
	}
	var input Blogs.CreateCommentRequest
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid type of input", err)
		return
	}

	createCommentData := Blogs.CommentModel{
		PostId:        postId,
		CommentedById: userId,
		ReplyPostId:   input.ReplyPostId,
		Comment:       input.Comment,
	}

	id, err := h.services.CreateComment(createCommentData)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "Error while creating post", err)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type getAllCommentsResponse struct {
	Data []Blogs.CommentResponse `json:"data"`
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
// @Success 			200 {object} getAllCommentsResponse
// @Failure 			400,404 {object} errorResponse
// @Failure				500 {object} errorResponse
// @Failure 			default {object} errorResponse
// @Router 				/posts/{post_id}/comments [get]
func (h *Handler) getAllComments(c *gin.Context) {
	postId, err := primitive.ObjectIDFromHex(c.Param("post_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid id param", err)
		return
	}
	comments, err := h.services.GetAllComments(postId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "error while getting comments", err)
		return
	}

	c.JSON(http.StatusOK, getAllCommentsResponse{
		Data: comments,
	})
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
// @Success 			200 {object} Blogs.CommentResponse
// @Failure 			400,404 {object} errorResponse
// @Failure				500 {object} errorResponse
// @Failure 			default {object} errorResponse
// @Router 				/posts/comments/{comment_id} [get]
func (h *Handler) getCommentById(c *gin.Context) {
	commentId, err := primitive.ObjectIDFromHex(c.Param("comment_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid comment_id param", err)
		return
	}

	comment, err := h.services.GetCommentById(commentId)
	if err != nil {
		if err.Error() == "no comments exist" {
			newErrorResponse(c, http.StatusInternalServerError, "No comments found", err)
			return
		}
		newErrorResponse(c, http.StatusInternalServerError, "Error in getting comment", err)
		return
	}

	c.JSON(http.StatusOK, comment)
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
// @Success 			200 {object} statusResponse
// @Failure 			400,404 {object} errorResponse
// @Failure				500 {object} errorResponse
// @Failure 			default {object} errorResponse
// @Router 				/posts/comments/{comment_id} [put]
func (h *Handler) updateComment(c *gin.Context) {
	commentId, err := primitive.ObjectIDFromHex(c.Param("comment_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid comment_id param", err)
		return
	}

	var input Blogs.UpdateCommentRequest
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid type of input", err)
		return
	}

	err = h.services.UpdateComment(commentId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "Error while updating comment", err)
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
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
// @Success 			200 {object} statusResponse
// @Failure 			400,404 {object} errorResponse
// @Failure				500 {object} errorResponse
// @Failure 			default {object} errorResponse
// @Router 				/posts/comments/{comment_id} [delete]
func (h *Handler) deleteComment(c *gin.Context) {
	commentId, err := primitive.ObjectIDFromHex(c.Param("comment_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid id param", err)
		return
	}

	err = h.services.DeleteComment(commentId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "Error while deleting comment", err)
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
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
// @Success 			200 {object} statusResponse
// @Failure 			400,404 {object} errorResponse
// @Failure				500 {object} errorResponse
// @Failure 			default {object} errorResponse
// @Router 				/posts/{post_id}/like [post]
func (h *Handler) likePost(c *gin.Context) {
	userId, err := h.getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "id not found", err)
		return
	}
	postId, err := primitive.ObjectIDFromHex(c.Param("post_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid id param", err)
		return
	}

	err = h.services.AddLike(postId, userId)
	if err != nil {
		if err.Error() == "already liked" {
			newErrorResponse(c, http.StatusInternalServerError, "Already liked", err)
			return
		}
		newErrorResponse(c, http.StatusInternalServerError, "Error while liking post", err)
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
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
// @Success 			200 {object} statusResponse
// @Failure 			400,404 {object} errorResponse
// @Failure				500 {object} errorResponse
// @Failure 			default {object} errorResponse
// @Router 				/posts/{post_id}/unlike [post]
func (h *Handler) unlikePost(c *gin.Context) {
	userId, err := h.getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "id not found", err)
		return
	}
	postId, err := primitive.ObjectIDFromHex(c.Param("post_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid id param", err)
		return
	}

	err = h.services.UnlikePost(postId, userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "Error while unliking post", err)
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}
