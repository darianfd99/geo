package handler

import (
	"errors"
	"log"
	"net/http"

	localization "github.com/darianfd99/geo/pkg"
	"github.com/darianfd99/geo/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.POST("/", h.postLocalization)
	router.GET("/", h.getAllLocalization)

	return router
}

type LocalizationRequest struct {
	Id   string   `json:"id" binding:"required"`
	Lat  *float64 `json:"lat" binding:"required"`
	Long *float64 `json:"long" binding:"required"`
}

func (h *Handler) postLocalization(c *gin.Context) {

	var input LocalizationRequest
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	group, err := h.services.Localization.Post(input.Id, *input.Lat, *input.Long)
	if err != nil {
		if errors.Is(err, localization.ErrInvalidCourseID) || errors.Is(err, service.ErrUuidExists) {
			newErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}
		newErrorResponse(c, http.StatusInternalServerError, "")
		return
	}

	log.Printf("Create position: longitude :%.2f --- latitude :%.2f\n", *input.Lat, *input.Long)
	c.JSON(http.StatusCreated, PostResponse{
		Group: group,
	})
}

func (h *Handler) getAllLocalization(c *gin.Context) {
	listLocalization, err := h.services.Localization.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, GetAllResponse{
		Data: listLocalization,
	})

}
