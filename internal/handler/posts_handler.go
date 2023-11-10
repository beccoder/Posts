package handler

import (
	"Blogs"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

func (h *Handler) createPost(c *gin.Context) {
	userId, err := h.getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "id not found", err)
		return
	}
	var input Blogs.Post
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid type of input", err)
		return
	}

	input.AuthorsId = userId
	id, err := h.services.CreatePosts(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "Error while creating post", err)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type getAllPostsResponse struct {
	Data []Blogs.Post `json:"data"`
}

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

func (h *Handler) updatePost(c *gin.Context) {
	postId, err := primitive.ObjectIDFromHex(c.Param("post_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid id param", err)
		return
	}

	var input Blogs.PostUpdate
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
	var input Blogs.Comment
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid type of input", err)
		return
	}

	input.CommentedById = userId
	input.PostId = postId
	id, err := h.services.CreateComment(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "Error while creating post", err)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type getAllCommentsResponse struct {
	Data []Blogs.Comment `json:"data"`
}

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

func (h *Handler) updateComment(c *gin.Context) {
	commentId, err := primitive.ObjectIDFromHex(c.Param("comment_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid comment_id param", err)
		return
	}

	var input Blogs.CommentUpdate
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

func (h *Handler) deleteComment(c *gin.Context) {}

func (h *Handler) createLike(c *gin.Context) {}

func (h *Handler) getAllLikes(c *gin.Context) {}

func (h *Handler) getLikeById(c *gin.Context) {}

func (h *Handler) deleteLike(c *gin.Context) {}
