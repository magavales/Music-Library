package handler

import "github.com/gin-gonic/gin"

// @Summary      Get song
// @Description  get song by id
// @Tags         get song
// @Produce      json
// @Param        id   path  integer  true  "The song's id"
// @Success      200  {object}  	models.Song
// @Failure      400  {object}  	error
// @Failure      500  {object}  	error
// @Router       /library/:id 	[get]
func (h *Handler) getSongByID(context *gin.Context) {
	h.Services.GetByID(context)
}
