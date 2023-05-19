package c_produk

import (
	"net/http"
	"toko-buah/service/s_produk"

	"github.com/gin-gonic/gin"
)

type ProdukController struct {
	getProdukService s_produk.GetProdukService
}

func NewGetProdukController(getProdukService s_produk.GetProdukService) *ProdukController {
	return &ProdukController{getProdukService}
}

func (c *ProdukController) GetAllProduk(ctx *gin.Context) {
	produk, err := c.getProdukService.GetAllProduk()
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	count := len(produk) // Calculate the count of the products

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Hasil data produk",
		"data":    produk,
		"count":   count, // Add the count to the response
	})
}
