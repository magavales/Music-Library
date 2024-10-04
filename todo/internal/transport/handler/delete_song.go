package handler

import "github.com/gin-gonic/gin"

func (h *Handler) deleteSong(context *gin.Context) {
	h.Services.Delete(context)
}
