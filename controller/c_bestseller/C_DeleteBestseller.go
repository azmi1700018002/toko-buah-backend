package c_bestseller

import (
	"net/http"
	"strconv"
	s_bestseller "toko-buah/service/s_best_seller"

	"github.com/gin-gonic/gin"
)

type deleteBestsellerController struct {
	deleteBestsellerService s_bestseller.DeleteBestsellerService
}

func NewBestsellerDeleteController(deleteBestsellerService s_bestseller.DeleteBestsellerService) *deleteBestsellerController {
	return &deleteBestsellerController{deleteBestsellerService}
}

func (c *deleteBestsellerController) DeleteBestseller(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("bestseller_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	err = c.deleteBestsellerService.DeleteBestsellerByID(id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": "Bestseller berhasil dihapus"})
}
