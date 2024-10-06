package handler

import "github.com/gin-gonic/gin"

// @Summary      Get songs
// @Description  get songs using the query's parameters for paginating and filtering
// @Tags         get songs
// @Produce      json
// @Param        limit   query  integer  true  "The query's parameters for paginating and filtering"
// @Param        offset   query  integer true  "The query's parameters for paginating and filtering"
// @Param        id   query  integer  true  "The query's parameters for paginating and filtering"
// @Param        group_name   query  string  true  "The query's parameters for paginating and filtering"
// @Param        song_name   query  string  true  "The query's parameters for paginating and filtering"
// @Param        release_date   query  string  true  "The query's parameters for paginating and filtering"
// @Param        text   query  string  true  "The query's parameters for paginating and filtering"
// @Param        link   query  string  true  "The query's parameters for paginating and filtering"
// @Success      200  {object}  	[]models.Song
// @Failure      400  {object}  	error
// @Failure      500  {object}  	error
// @Router       /library 	[get]
func (h *Handler) getLibrary(context *gin.Context) {
	h.Services.Get(context)
}
