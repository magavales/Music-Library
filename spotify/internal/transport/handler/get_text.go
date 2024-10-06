package handler

import "github.com/gin-gonic/gin"

// @Summary      Get the song's text
// @Description  get the song's text using the query's parameters for paginating and filtering
// @Tags         get the song's text
// @Produce      json
// @Param        id   path  integer  true  "The query's parameters for paginating and filtering"
// @Success      200  {string}  	string
// @Failure      400  {object}  	error
// @Failure      500  {object}  	error
// @Router       /library/:id/text 	[get]
func (h *Handler) getText(context *gin.Context) {
	h.Services.GetText(context)
}
