package handler

import (
	"log"

	localization "github.com/darianfd99/geo/pkg"
	"github.com/gin-gonic/gin"
)

const internalError = "Internal Error"

type errorResponse struct {
	Message string `json:"message"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	log.Printf("error :%s", message)
	if statusCode == 500 {
		c.AbortWithStatusJSON(statusCode, errorResponse{internalError})
		return
	}
	c.AbortWithStatusJSON(statusCode, errorResponse{message})
}

type StatusResponse struct {
	Status string `json:"status"`
}

type PostResponse struct {
	Group localization.GroupAvaerageLocation `json:"group"`
}

type GetAllResponse struct {
	Data interface{} `json:"data"`
}
