package handler

import (
	"Blogs"
	"github.com/gin-gonic/gin"
	"net/http"
)

// signUpAuthor godoc
// @Summary			Sign Up Author
// @Tags 			Authorization
// @Description		Create account for author
// @ID 				create-account-author
// @Accept 			json
// @Produce 		json
// @Param 			input body Blogs.User true "account info"
// @Success 		200 {integer} integer 1
// @Failure 		400,404 {object} errorResponse
// @Failure			500 {object} errorResponse
// @Failure 		default {object} errorResponse
// @Router 			/auth/author/sign-up [post]
func (h *Handler) signUpAuthor(c *gin.Context) {
	var input Blogs.User
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Error with input", err)
		return
	}
	input.Role = "author"
	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		if err.Error() == "already registered" {
			newErrorResponse(c, http.StatusInternalServerError, "Author with provided data is already registered", err)
			return
		}
		if err.Error() == "username exists" {
			newErrorResponse(c, http.StatusInternalServerError, "Username exists", err)
			return
		}
		newErrorResponse(c, http.StatusInternalServerError, "Error while creating author", err) //Recheck
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

// signUpUser godoc
// @Summary			Sign Up User
// @Tags 			Authorization
// @Description		Create account for user
// @ID 				create-account-user
// @Accept 			json
// @Produce 		json
// @Param 			input body Blogs.User true "account info"
// @Success 		200 {integer} integer 1
// @Failure 		400,404 {object} errorResponse
// @Failure			500 {object} errorResponse
// @Failure 		default {object} errorResponse
// @Router 			/auth/user/sign-up [post]
func (h *Handler) signUpUser(c *gin.Context) {
	var input Blogs.User
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Error with user input", err)
		return
	}
	input.Role = "user"
	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		if err.Error() == "already registered" {
			newErrorResponse(c, http.StatusInternalServerError, "User with provided data is already registered", err)
			return
		}
		if err.Error() == "username exists" {
			newErrorResponse(c, http.StatusInternalServerError, "Username exists", err)
			return
		}

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

// signInAuthor godoc
// @Summary			Sign In Author
// @Tags 			Authorization
// @Description		Login to author account
// @ID 				login-author-account
// @Accept 			json
// @Produce 		json
// @Param 			input body signInInput true "credentials"
// @Success 		200 {object} TokenRoleResponse
// @Failure 		400,404 {object} errorResponse
// @Failure			500 {object} errorResponse
// @Failure 		default {object} errorResponse
// @Router 			/auth/author/sign-in [post]
func (h *Handler) signInAuthor(c *gin.Context) {
	var input signInInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid type of input", err)
		return
	}

	token, err := h.services.GenerateToken(input.Username, input.Password, "author")
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

// signInUser godoc
// @Summary			Sign In User
// @Tags 			Authorization
// @Description		Login to user account
// @ID 				login-user-account
// @Accept 			json
// @Produce 		json
// @Param 			input body signInInput true "credentials"
// @Success 		200 {object} TokenRoleResponse
// @Failure 		400,404 {object} errorResponse
// @Failure			500 {object} errorResponse
// @Failure 		default {object} errorResponse
// @Router 			/auth/user/sign-in [post]
func (h *Handler) signInUser(c *gin.Context) {
	var input signInInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid type of input", err)
		return
	}

	token, err := h.services.GenerateToken(input.Username, input.Password, "user")
	if err != nil {
		if err.Error() == "Invalid role" {
			newErrorResponse(c, http.StatusInternalServerError, "Invalid role", err)
			return
		}
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
