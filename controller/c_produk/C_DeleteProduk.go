package c_produk

import (
	"net/http"
	"strconv"
	"toko-buah/service/s_produk"

	"github.com/gin-gonic/gin"
)

type deleteProdukController struct {
	deleteProdukService s_produk.DeleteProdukService
}

func NewProdukDeleteController(deleteProdukService s_produk.DeleteProdukService) *deleteProdukController {
	return &deleteProdukController{deleteProdukService}
}

func (c *deleteProdukController) DeleteProduk(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("produk_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	err = c.deleteProdukService.DeleteProdukByID(id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": "Produk berhasil dihapus"})
}
