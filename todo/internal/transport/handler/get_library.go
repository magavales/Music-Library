package handler

import "github.com/gin-gonic/gin"

func (h *Handler) getLibrary(context *gin.Context) {
	h.Services.Get(context)
}
