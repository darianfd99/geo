package handler

import (
	"fmt"

	localization "github.com/darianfd99/geo/pkg"
	"github.com/gin-gonic/gin"
)

type errorResponse struct {
	Message string `json:"message"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	fmt.Printf("error %d: %s", statusCode, message)
	c.AbortWithStatusJSON(statusCode, errorResponse{message})
}

type StatusResponse struct {
	Status string `json:"status"`
}

type GetAllLocalizationResponse struct {
	Data []localization.Localization
}
