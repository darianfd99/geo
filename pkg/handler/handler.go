package handler

import (
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

func (h *Handler) postLocalization(c *gin.Context) {

	var input localization.Localization
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err := h.services.Localization.Post(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, StatusResponse{
		Status: "OK",
	})
}

func (h *Handler) getAllLocalization(c *gin.Context) {
	listLocalization, err := h.services.Localization.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, GetAllLocalizationResponse{
		Data: listLocalization,
	})

}
