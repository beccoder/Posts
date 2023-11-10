package handler

import (
	"Blogs"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) signUpAuthor(c *gin.Context) {
	var input Blogs.User
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Error with user input", err)
		return
	}
	input.Role = "author"
	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "Error while creating author", err) //Recheck
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) signUpUser(c *gin.Context) {
	var input Blogs.User
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Error with user input", err)
		return
	}
	input.Role = "user"
	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "Error while creating user", err) //Recheck
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type signInInput struct {
	Username string `json:"username" bson:"username" binding:"required"`
	Password string `json:"password" bson:"password" binding:"required"`
}

func (h *Handler) signInAuthor(c *gin.Context) {
	var input signInInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid type of input", err)
		return
	}

	token, err := h.services.Authorization.GenerateToken(input.Username, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "Invalid username or password", err)
		return
	}

	role, err := h.services.GetUserRole(input.Username, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "Error with user role", err)
		return
	}

	if role != "author" {
		newErrorResponse(c, http.StatusInternalServerError, "User does not have author access", err)
		return
	}

	c.JSON(http.StatusOK, TokenRoleResponse{
		Token: token,
		Role:  role,
	})
}

func (h *Handler) signInUser(c *gin.Context) {
	var input signInInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid type of input", err)
		return
	}

	token, err := h.services.Authorization.GenerateToken(input.Username, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "Invalid username or password", err)
		return
	}

	role, err := h.services.GetUserRole(input.Username, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "Error with user role", err)
		return
	}

	if role != "user" {
		newErrorResponse(c, http.StatusInternalServerError, "User has author access", err)
		return
	}

	c.JSON(http.StatusOK, TokenRoleResponse{
		Token: token,
		Role:  role,
	})
}
