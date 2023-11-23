package http

import (
	"github.com/gin-gonic/gin"
	"log"
)

type Response struct {
	Status      string      `json:"status"`
	Description string      `json:"description"`
	Data        interface{} `json:"data"`
}

func HandleResponse(c *gin.Context, status Status, data interface{}) {
	switch code := status.Code; {
	case code < 300:
		log.Printf("---Response--->\ncode: %d\nstatus: %s\ndescription: %s\n", status.Code, status.Status, status.Description)
		log.Println("data: ", data)

	case code < 400:
		log.Printf("---Warn--->\ncode: %d\nstatus: %s\ndescription: %s\n", status.Code, status.Status, status.Description)
		log.Println("data: ", data)

	default:
		log.Printf("---Error--->\ncode: %d\nstatus: %s\ndescription: %s\n", status.Code, status.Status, status.Description)
		log.Println("data: ", data)
	}

	c.JSON(status.Code, Response{
		Status:      status.Status,
		Description: status.Description,
		Data:        data,
	})
}
