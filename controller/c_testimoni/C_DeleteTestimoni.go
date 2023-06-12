package c_testimoni

import (
	"net/http"
	"strconv"
	s_testimoni "toko-buah/service/s_testimoni"

	"github.com/gin-gonic/gin"
)

type deleteTestimoniController struct {
	deleteTestimoniService s_testimoni.DeleteTestimoniService
}

func NewTestimoniDeleteController(deleteTestimoniService s_testimoni.DeleteTestimoniService) *deleteTestimoniController {
	return &deleteTestimoniController{deleteTestimoniService}
}

func (c *deleteTestimoniController) DeleteTestimoni(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("testimoni_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	err = c.deleteTestimoniService.DeleteTestimoniByID(id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": "Testimoni berhasil dihapus"})
}
