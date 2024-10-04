package handler

import "github.com/gin-gonic/gin"

func (h *Handler) createSong(context *gin.Context) {
	h.Services.Create(context)
}
