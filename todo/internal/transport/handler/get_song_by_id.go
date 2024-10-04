package handler

import "github.com/gin-gonic/gin"

func (h *Handler) getSongByID(context *gin.Context) {
	h.Services.GetByID(context)
}
