package handler

import (
	"github.com/gin-gonic/gin"
	"log"
)

type errorResponse struct {
	Message string `json:"message"`
}

type statusResponse struct {
	Status string `json:"status"`
}

type TokenRoleResponse struct {
	Token string `json:"token"`
	Role  string `json:"role"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string, err error) {
	log.Printf(message+": %s", err.Error())
	c.AbortWithStatusJSON(statusCode, errorResponse{message})
}
