package handler

import "github.com/gin-gonic/gin"

// @Summary      Delete song
// @Description  Delete song
// @Tags         Delete
// @Produce      json
// @Param        id   path  integer  true  "The song's id"
// @Success      200  {string}  	string "deleted"
// @Failure      400  {object}  	error
// @Failure      500  {object}  	error
// @Router       /library/:id 	[delete]
func (h *Handler) deleteSong(context *gin.Context) {
	h.Services.Delete(context)
}
