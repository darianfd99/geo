package handler

import (
	localization "github.com/darianfd99/geo/pkg"
	"github.com/gin-gonic/gin"
)

type errorResponse struct {
	Message string `json:"message"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
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
