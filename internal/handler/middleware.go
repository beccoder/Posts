package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	userCtxId           = "userId"
	userCtxRole         = "role"
)

func (h *Handler) middlewareAdmin(c *gin.Context) {}

func (h *Handler) middlewareAuthor(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		newErrorResponse(c, http.StatusUnauthorized, "Empty auth header", errors.New("Empty auth header"))
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		newErrorResponse(c, http.StatusUnauthorized, "Invalid auth header", errors.New("Invalid auth header"))
		return
	}

	userId, role, err := h.services.Authorization.ParseToken(headerParts[1])
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, "Token error", err)
		return
	}
	if role != "author" {
		newErrorResponse(c, http.StatusUnauthorized, "User is not author", errors.New("User is not author"))
		return

	}
	c.Set(userCtxId, userId)
	c.Set(userCtxRole, role)
}

func (h *Handler) middleware(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		newErrorResponse(c, http.StatusUnauthorized, "Empty auth header", errors.New("Empty auth header"))
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		newErrorResponse(c, http.StatusUnauthorized, "Invalid auth header", errors.New("Invalid auth header"))
		return
	}

	userId, role, err := h.services.Authorization.ParseToken(headerParts[1])
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, "Token error", err)
		return
	}
	c.Set(userCtxId, userId)
	c.Set(userCtxRole, role)
}

func (h *Handler) getUserId(c *gin.Context) (primitive.ObjectID, error) {
	id, ok := c.Get(userCtxId)
	if !ok {
		return primitive.ObjectID{}, errors.New("id not found")
	}

	idTypeObjectId, ok := id.(primitive.ObjectID)
	if !ok {
		return primitive.ObjectID{}, errors.New("id not found")
	}

	return idTypeObjectId, nil
}

func getUserRole(c *gin.Context) (string, error) {
	role, ok := c.Get(userCtxRole)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "role not found", errors.New("role not found"))
		return "", errors.New("role not found")
	}
	roleStr, ok := role.(string)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "role is invalid of type", errors.New("role is invalid of type"))
		return "", errors.New("role not found")
	}
	return roleStr, nil
}

func (h *Handler) checkOwnershipComment(c *gin.Context) {
	commentId, err := primitive.ObjectIDFromHex(c.Param("comment_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid comment_id param", err)
		return
	}
	userId, err := h.getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "id not found", err)
		return
	}

	// check for ownership
	comment, err := h.services.GetCommentById(commentId)
	if err != nil {
		if err.Error() == "no comments exist" {
			newErrorResponse(c, http.StatusInternalServerError, "No comments found", err)
			return
		}
		newErrorResponse(c, http.StatusInternalServerError, "Error in getting comment", err)
		return
	}
	if userId.Hex() != comment.CommentedById.Hex() {
		newErrorResponse(c, http.StatusInternalServerError, "No access to update or delete comment", errors.New("no access"))
		return
	}
}

func (h *Handler) checkOwnershipPost(c *gin.Context) {
	postId, err := primitive.ObjectIDFromHex(c.Param("post_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid comment_id param", err)
		return
	}

	userId, err := h.getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "id not found", err)
		return
	}

	// check for ownership
	post, err := h.services.GetPostById(postId)
	if err != nil {
		if err.Error() == "no posts exist" {
			newErrorResponse(c, http.StatusInternalServerError, "No comments found", err)
			return
		}
		newErrorResponse(c, http.StatusInternalServerError, "Error in getting comment", err)
		return
	}

	if userId.Hex() != post.AuthorsId.Hex() {
		newErrorResponse(c, http.StatusInternalServerError, "No access to update or delete comment", errors.New("no access"))
		return
	}
}
