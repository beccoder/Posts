package handler

import (
	"Blogs"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

type getAllUsersResponse struct {
	Data []Blogs.UserResponse `json:"data"`
}

func (h *Handler) getAllUsers(c *gin.Context) {
	users, err := h.services.GetAllUsers()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "error while getting users", err)
		return
	}

	c.JSON(http.StatusOK, getAllUsersResponse{
		Data: users,
	})
}

func (h *Handler) getUserById(c *gin.Context) {
	userId, err := primitive.ObjectIDFromHex(c.Param("user_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid user_id param", err)
		return
	}

	user, err := h.services.GetUserById(userId)
	if err != nil {
		if err.Error() == "no matching user" {
			newErrorResponse(c, http.StatusNotFound, "No users found", err)
			return
		}
		newErrorResponse(c, http.StatusInternalServerError, "Error in getting user", err)
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *Handler) updateUser(c *gin.Context) {
	userId, err := primitive.ObjectIDFromHex(c.Param("user_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid user_id param", err)
		return
	}

	var input Blogs.UpdateUserRequest
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid type of input", err)
		return
	}

	err = h.services.UpdateUser(userId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "Error while updating user", err)
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}

func (h *Handler) deleteUser(c *gin.Context) {
	userId, err := primitive.ObjectIDFromHex(c.Param("user_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid id param", err)
		return
	}

	err = h.services.DeleteUser(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "Error while deleting user", err)
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}
