package handler

import (
	"github.com/gin-gonic/gin"
)

// @Summary      Update song
// @Description  Update song
// @Tags         Update
// @Accept       json
// @Produce      json
// @Param		 id path integer true "The song's id"
// @Param        song   body  models.Song  true  "The song's info"
// @Success      200  {string}  	string "updated"
// @Failure      400  {object}  	error
// @Failure      500  {object}  	error
// @Router       /library/:id 	[patch]
func (h *Handler) updateSong(context *gin.Context) {
	h.Services.Update(context)
}
