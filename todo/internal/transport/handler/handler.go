package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/magavales/Music-Library/todo/internal/services"
)

type Handler struct {
	Services *services.Service
}

func NewHandler(services *services.Service) *Handler {
	return &Handler{
		Services: services,
	}
}

func (h *Handler) InitRouter() *gin.Engine {
	router := gin.Default()

	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			library := v1.Group("/library")
			{
				library.GET("/:id", h.getSongByID)
				library.GET("", h.getLibrary)
				library.POST("", h.createSong)
				library.DELETE("/:id", h.deleteSong)
				library.PATCH("/:id", h.updateSong)
				library.GET("/:id/text", h.getText)
			}
		}
	}

	return router
}
