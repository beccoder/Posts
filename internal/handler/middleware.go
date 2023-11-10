package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	userCtxId           = "userId"
	userCtxRole         = "role"
)

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
		newErrorResponse(c, http.StatusUnauthorized, "User is not admin", errors.New("User is not admin"))
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
