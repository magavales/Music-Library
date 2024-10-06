package handler

import (
	"github.com/gin-gonic/gin"
)

// @Summary      Create song
// @Description  create song
// @Tags         create
// @Accept       json
// @Produce      json
// @Param        song   body  models.Song  true  "The song's info"
// @Success      200  {integer}  	integer 1
// @Failure      400  {object}  	error
// @Failure      500  {object}  	error
// @Router       /library 	[post]
func (h *Handler) createSong(context *gin.Context) {
	h.Services.Create(context)
}
