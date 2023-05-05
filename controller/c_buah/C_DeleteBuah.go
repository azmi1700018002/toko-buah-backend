package c_buah

import (
	"net/http"
	"strconv"
	"toko-buah/service/s_buah"

	"github.com/gin-gonic/gin"
)

type deleteBuahController struct {
	deleteBuahService s_buah.DeleteBuahService
}

func NewBuahDeleteController(deleteBuahService s_buah.DeleteBuahService) *deleteBuahController {
	return &deleteBuahController{deleteBuahService}
}

func (c *deleteBuahController) DeleteBuah(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("buah_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	err = c.deleteBuahService.DeleteBuahByID(id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": "Buah berhasil dihapus"})
}
