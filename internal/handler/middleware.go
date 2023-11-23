package handler

import (
	"Blogs/internal/handler/http"
	"errors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	userCtxId           = "userId"
	userCtxRole         = "role"
)

func (h *Handler) middlewareAdmin(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		http.HandleResponse(c, http.Unauthorized, "Empty auth header")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		http.HandleResponse(c, http.InvalidAuthHeader, "Invalid auth header")
		return
	}

	userId, role, err := h.services.Authorization.ParseToken(headerParts[1])
	if err != nil {
		http.HandleResponse(c, http.InvalidAuth, err.Error())
		return
	}

	if role != "admin" {
		http.HandleResponse(c, http.Forbidden, "Forbidden")
		return

	}
	c.Set(userCtxId, userId)
	c.Set(userCtxRole, role)
}

func (h *Handler) middlewareAuthor(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		http.HandleResponse(c, http.Unauthorized, "Empty auth header")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		http.HandleResponse(c, http.InvalidAuthHeader, "Invalid auth header")
		return
	}

	userId, role, err := h.services.Authorization.ParseToken(headerParts[1])
	if err != nil {
		http.HandleResponse(c, http.InvalidAuth, err.Error())
		return
	}
	if role != "author" && role != "admin" {
		http.HandleResponse(c, http.AccessDenied, "Access denied")
		return

	}
	c.Set(userCtxId, userId)
	c.Set(userCtxRole, role)
}

func (h *Handler) middleware(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		http.HandleResponse(c, http.Unauthorized, "Empty auth header")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		http.HandleResponse(c, http.InvalidAuthHeader, "Invalid auth header")
		return
	}

	userId, role, err := h.services.Authorization.ParseToken(headerParts[1])
	if err != nil {
		http.HandleResponse(c, http.InvalidAuth, err.Error())
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
		return "", errors.New("role not found")
	}
	roleStr, ok := role.(string)
	if !ok {
		return "", errors.New("role not string")
	}
	return roleStr, nil
}

func (h *Handler) checkOwnershipComment(c *gin.Context) {
	commentId, err := primitive.ObjectIDFromHex(c.Param("comment_id"))
	if err != nil {
		http.HandleResponse(c, http.InvalidArgument, err.Error())
		return
	}
	userId, err := h.getUserId(c)
	if err != nil {
		http.HandleResponse(c, http.InternalServerError, err.Error())
		return
	}

	// check for ownership
	comment, err := h.services.GetCommentById(commentId)
	if err != nil {
		if err.Error() == "no comments exist" {
			http.HandleResponse(c, http.NotFound, err.Error())
			return
		}
		http.HandleResponse(c, http.InternalServerError, err.Error())
		return
	}
	role, err := getUserRole(c)
	if err != nil {
		http.HandleResponse(c, http.InvalidArgument, err.Error())
		return
	}
	if userId.Hex() != comment.CommentedById.Hex() && role != "admin" {
		http.HandleResponse(c, http.AccessDenied, "access denied")
		return
	}
}

func (h *Handler) checkOwnershipPost(c *gin.Context) {
	postId, err := primitive.ObjectIDFromHex(c.Param("post_id"))
	if err != nil {
		http.HandleResponse(c, http.InvalidArgument, err.Error())
		return
	}

	userId, err := h.getUserId(c)
	if err != nil {
		http.HandleResponse(c, http.InvalidArgument, err.Error())
		return
	}

	// check for ownership
	post, err := h.services.GetPostById(postId)
	if err != nil {
		if err.Error() == "no posts exist" {
			http.HandleResponse(c, http.NotFound, err.Error())
			return
		}
		http.HandleResponse(c, http.InternalServerError, err.Error())
		return
	}

	role, err := getUserRole(c)
	if err != nil {
		http.HandleResponse(c, http.InvalidArgument, err.Error())
		return
	}

	if userId.Hex() != post.AuthorsId.Hex() && role != "admin" {
		http.HandleResponse(c, http.AccessDenied, "access denied")
		return
	}
}
