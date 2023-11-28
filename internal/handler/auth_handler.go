package handler

import (
	"Blogs"
	"Blogs/internal/handler/http"
	"github.com/gin-gonic/gin"
	"strings"
)

// signUpAuthor godoc
// @Summary			Sign Up Author
// @Tags 			Authorization
// @Description		Create account for author
// @ID 				create-account-author
// @Accept 			json
// @Produce 		json
// @Param 			input body Blogs.SignUpUserRequest true "account info"
// @Success 		200 {object} http.Response ""
// @Failure 		400,404 {object} http.ErrorResponse
// @Failure			500 {object} http.ErrorResponse
// @Failure 		default {object} http.ErrorResponse
// @Router 			/auth/author/sign-up [post]
func (h *Handler) signUpAuthor(c *gin.Context) {
	var input Blogs.SignUpUserRequest
	if err := c.BindJSON(&input); err != nil {
		http.HandleErrorResponse(c, http.BadRequest, err.Error())
		return
	}
	signUpData := Blogs.UserModel{
		Role:      "author",
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Username:  strings.ToLower(input.Username),
		Password:  input.Password,
		Email:     strings.ToLower(input.Email),
		Bio:       input.Bio,
	}
	id, err := h.services.CreateUser(signUpData)
	if err != nil {
		if err.Error() == "already registered" {
			http.HandleErrorResponse(c, http.RequestConflict, err.Error())
			return
		}
		if err.Error() == "username exists" {
			http.HandleErrorResponse(c, http.RequestConflict, err.Error())
			return
		}
		http.HandleErrorResponse(c, http.InternalServerError, err.Error())
		return
	}

	http.HandleResponse(c, http.OK, map[string]interface{}{
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
// @Param 			input body Blogs.SignUpUserRequest true "account info"
// @Success 		200 {object} http.Response ""
// @Failure 		400,404 {object} http.ErrorResponse
// @Failure			500 {object} http.ErrorResponse
// @Failure 		default {object} http.ErrorResponse
// @Router 			/auth/user/sign-up [post]
func (h *Handler) signUpUser(c *gin.Context) {
	var input Blogs.SignUpUserRequest
	if err := c.BindJSON(&input); err != nil {
		http.HandleErrorResponse(c, http.BadRequest, err.Error())
		return
	}
	signUpData := Blogs.UserModel{
		Role:      "user",
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Username:  strings.ToLower(input.Username),
		Password:  input.Password,
		Email:     strings.ToLower(input.Email),
		Bio:       input.Bio,
	}
	id, err := h.services.CreateUser(signUpData)
	if err != nil {
		if err.Error() == "already registered" {
			http.HandleErrorResponse(c, http.RequestConflict, err.Error())
			return
		}
		if err.Error() == "username exists" {
			http.HandleErrorResponse(c, http.RequestConflict, err.Error())
			return
		}
		http.HandleErrorResponse(c, http.InternalServerError, err.Error())
		return
	}

	http.HandleResponse(c, http.OK, map[string]interface{}{
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
// @Success 		200 {object} http.Response
// @Failure 		400,404 {object} http.ErrorResponse
// @Failure			500 {object} http.ErrorResponse
// @Failure 		default {object} http.ErrorResponse
// @Router 			/auth/author/sign-in [post]
func (h *Handler) signInAuthor(c *gin.Context) {
	var input signInInput
	if err := c.BindJSON(&input); err != nil {
		http.HandleErrorResponse(c, http.BadRequest, err.Error())
		return
	}

	token, err := h.services.GenerateToken(strings.ToLower(input.Username), input.Password, "author")
	if err != nil {
		if err.Error() == "Invalid role" {
			http.HandleErrorResponse(c, http.InvalidArgument, err.Error())
			return
		}
		http.HandleErrorResponse(c, http.InternalServerError, err.Error())
		return
	}

	http.HandleResponse(c, http.OK, TokenRoleResponse{
		Token: token,
		Role:  "author",
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
// @Success 		200 {object} http.Response
// @Failure 		400,404 {object} http.ErrorResponse
// @Failure			500 {object} http.ErrorResponse
// @Failure 		default {object} http.ErrorResponse
// @Router 			/auth/user/sign-in [post]
func (h *Handler) signInUser(c *gin.Context) {
	var input signInInput
	if err := c.BindJSON(&input); err != nil {
		http.HandleErrorResponse(c, http.BadRequest, err.Error())
		return
	}

	token, err := h.services.GenerateToken(strings.ToLower(input.Username), input.Password, "user")
	if err != nil {
		if err.Error() == "Invalid role" {
			http.HandleErrorResponse(c, http.InvalidArgument, err.Error())
			return
		}
		http.HandleErrorResponse(c, http.InternalServerError, err.Error())
		return
	}

	http.HandleResponse(c, http.OK, TokenRoleResponse{
		Token: token,
		Role:  "user",
	})
}

// signInAdmin godoc
// @Summary			Sign In Admin
// @Tags 			Authorization
// @Description		Login to admin account
// @ID 				login-admin-account
// @Accept 			json
// @Produce 		json
// @Param 			input body signInInput true "credentials"
// @Success 		200 {object} http.Response
// @Failure 		400,404 {object} http.ErrorResponse
// @Failure			500 {object} http.ErrorResponse
// @Failure 		default {object} http.ErrorResponse
// @Router 			/auth/admin/sign-in [post]
func (h *Handler) signInAdmin(c *gin.Context) {
	var input signInInput
	if err := c.BindJSON(&input); err != nil {
		http.HandleErrorResponse(c, http.BadRequest, err.Error())
		return
	}

	token, err := h.services.GenerateToken(strings.ToLower(input.Username), input.Password, "admin")
	if err != nil {
		if err.Error() == "Invalid role" {
			http.HandleErrorResponse(c, http.InvalidArgument, err.Error())
			return
		}
		http.HandleErrorResponse(c, http.InternalServerError, err.Error())
		return
	}

	http.HandleResponse(c, http.OK, TokenRoleResponse{
		Token: token,
		Role:  "admin",
	})
}
