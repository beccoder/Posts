package handler

import (
	"Blogs"
	"Blogs/internal/handler/http"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// createUser godoc
// @Summary			Create User
// @Security 			ApiKeyAuth
// @Tags 			Administration
// @Description		Create account
// @ID 				create-account
// @Accept 			json
// @Produce 		json
// @Param 			input body Blogs.CreateUserRequest true "account info"
// @Success 		200 {object} http.Response ""
// @Failure 		400,404 {object} http.Response
// @Failure			500 {object} http.Response
// @Failure 		default {object} http.Response
// @Router 			/admin/users [post]
func (h *Handler) createUser(c *gin.Context) {
	var input Blogs.CreateUserRequest
	if err := c.BindJSON(&input); err != nil {
		http.HandleResponse(c, http.BadRequest, err.Error())
		return
	}
	signUpData := Blogs.UserModel{
		Role:      input.Role,
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Username:  input.Username,
		Password:  input.Password,
		Email:     input.Email,
		Bio:       input.Bio,
	}
	id, err := h.services.CreateUser(signUpData)
	if err != nil {
		if err.Error() == "already registered" {
			http.HandleResponse(c, http.RequestConflict, err.Error())
			return
		}
		if err.Error() == "username exists" {
			http.HandleResponse(c, http.RequestConflict, err.Error())
			return
		}
		http.HandleResponse(c, http.InternalServerError, err.Error())
		return
	}

	http.HandleResponse(c, http.Created, map[string]interface{}{
		"id": id,
	})
}

// getAllUsers godoc
// @Summary 			Get All Users
// @Security 			ApiKeyAuth
// @Tags 				Administration
// @Description 		Get a list of all users
// @ID 					get-users
// @Accept 				json
// @Produce 			json
// @Success 			200 {object} http.Response
// @Failure 			400,404 {object} http.Response
// @Failure				500 {object} http.Response
// @Failure 			default {object} http.Response
// @Router 				/admin/users [get]
func (h *Handler) getAllUsers(c *gin.Context) {
	users, err := h.services.GetAllUsers()
	if err != nil {
		http.HandleResponse(c, http.InternalServerError, err.Error())
		return
	}
	http.HandleResponse(c, http.OK, users)
}

// getUserById godoc
// @Summary 			Get a User by ID
// @Security 			ApiKeyAuth
// @Tags 				Administration
// @Description 		Get a single User by providing its ID
// @ID 					get-user
// @Accept 				json
// @Produce 			json
// @Param 				users_id path string true "User ID"
// @Success 			200 {object} http.Response
// @Failure 			400,404 {object} http.Response
// @Failure				500 {object} http.Response
// @Failure 			default {object} http.Response
// @Router 				/admin/users/{users_id} [get]
func (h *Handler) getUserById(c *gin.Context) {
	userId, err := primitive.ObjectIDFromHex(c.Param("users_id"))
	if err != nil {
		http.HandleResponse(c, http.BadRequest, err.Error())
		return
	}

	user, err := h.services.GetUserById(userId)
	if err != nil {
		if err.Error() == "no matching user" {
			http.HandleResponse(c, http.NotFound, err.Error())
			return
		}
		http.HandleResponse(c, http.InternalServerError, err.Error())
		return
	}

	http.HandleResponse(c, http.OK, user)
}

// updateUser godoc
// @Summary 			Update an existing User data
// @Security 			ApiKeyAuth
// @Tags 				Administration
// @Description 		Update an existing User data with new data
// @ID 					update-user
// @Accept 				json
// @Produce 			json
// @Param 				users_id path string true "User ID"
// @Param 				new_input body Blogs.UpdateUserRequest true "new user info"
// @Success 			200 {object} http.Response
// @Failure 			400,404 {object} http.Response
// @Failure				500 {object} http.Response
// @Failure 			default {object} http.Response
// @Router 				/admin/users/{users_id} [put]
func (h *Handler) updateUser(c *gin.Context) {
	userId, err := primitive.ObjectIDFromHex(c.Param("users_id"))
	if err != nil {
		http.HandleResponse(c, http.BadRequest, err.Error())
		return
	}

	var input Blogs.UpdateUserRequest
	if err := c.BindJSON(&input); err != nil {
		http.HandleResponse(c, http.BadRequest, err.Error())
		return
	}

	err = h.services.UpdateUser(userId, input)
	if err != nil {
		http.HandleResponse(c, http.InternalServerError, err.Error())
		return
	}

	http.HandleResponse(c, http.OK, statusResponse{"ok"})
}

// deleteUser godoc
// @Summary				Delete an existing User
// @Security 			ApiKeyAuth
// @Tags 				Administration
// @Description			Delete an existing User by ID
// @ID 					delete-user
// @Accept 				json
// @Produce 			json
// @Param 				users_id path string true "User ID"
// @Success 			200 {object} http.Response
// @Failure 			400,404 {object} http.Response
// @Failure				500 {object} http.Response
// @Failure 			default {object} http.Response
// @Router 				/admin/users/{users_id} [delete]
func (h *Handler) deleteUser(c *gin.Context) {
	userId, err := primitive.ObjectIDFromHex(c.Param("users_id"))
	if err != nil {
		http.HandleResponse(c, http.BadRequest, err.Error())
		return
	}

	err = h.services.DeleteUser(userId)
	if err != nil {
		http.HandleResponse(c, http.InternalServerError, err.Error())
		return
	}

	http.HandleResponse(c, http.OK, statusResponse{"ok"})
}
