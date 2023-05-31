package c_produk

import (
	"net/http"
	"toko-buah/model/m_produk"
	"toko-buah/service/s_produk"

	"github.com/gin-gonic/gin"
)

type UpdateProdukController struct {
	updateProdukService *s_produk.UpdateProdukService
}

func NewUpdateProdukController(updateProdukService *s_produk.UpdateProdukService) *UpdateProdukController {
	return &UpdateProdukController{updateProdukService}
}

func (c *UpdateProdukController) UpdateProduk(ctx *gin.Context) {
	// Mendapatkan data produk dari request body
	var produk m_produk.Produk
	err := ctx.ShouldBind(&produk)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Mendapatkan file foto profil dari request body
	file, err := ctx.FormFile("gambar")
	if err != nil && err != http.ErrMissingFile {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Memanggil service untuk update produk
	updatedProduk, err := c.updateProdukService.UpdateProduk(ctx, produk, file)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Menampilkan respon JSON
	ctx.JSON(http.StatusOK, gin.H{"data": updatedProduk, "message": "Produk berhasil diperbarui"})
}
