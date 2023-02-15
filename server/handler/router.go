package handler

import (
	"github.com/gin-gonic/gin"

	"github.com/klimovI/go_tickers_rates/server/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) Init() *gin.Engine {
	router := gin.Default()

	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			v1.GET("/rates", h.getRates)
			v1.POST("/rates", h.postRates)
		}
	}

	return router
}
