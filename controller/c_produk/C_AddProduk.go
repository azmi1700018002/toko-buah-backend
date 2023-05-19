package c_produk

import (
	"net/http"
	"toko-buah/model/m_produk"
	"toko-buah/service/s_produk"

	"github.com/gin-gonic/gin"
)

type addProdukController struct {
	addProdukService *s_produk.AddProdukService
}

func NewProdukAddController(addProdukService *s_produk.AddProdukService) *addProdukController {
	return &addProdukController{addProdukService}
}

func (c *addProdukController) AddProduk(ctx *gin.Context) {
	var produk m_produk.Produk
	if err := ctx.ShouldBind(&produk); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	addProduk, err := c.addProdukService.AddProduk(ctx, produk)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": addProduk, "message": "Produk Telah Berhasil Ditambahkan"})
}
