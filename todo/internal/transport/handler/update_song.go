package handler

import (
	"github.com/gin-gonic/gin"
)

func (h *Handler) updateSong(context *gin.Context) {
	h.Services.Update(context)
}
