package handler

import "github.com/gin-gonic/gin"

func (h *Handler) getText(context *gin.Context) {
	h.Services.GetText(context)
}
