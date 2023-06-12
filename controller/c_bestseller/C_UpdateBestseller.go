package c_bestseller

import (
	"net/http"
	m_bestseller "toko-buah/model/m_best_seller"
	s_bestseller "toko-buah/service/s_best_seller"

	"github.com/gin-gonic/gin"
)

type UpdateBestsellerController struct {
	updateBestsellerService *s_bestseller.UpdateBestsellerService
}

func NewUpdateBestsellerController(updateBestsellerService *s_bestseller.UpdateBestsellerService) *UpdateBestsellerController {
	return &UpdateBestsellerController{updateBestsellerService}
}

func (c *UpdateBestsellerController) UpdateBestseller(ctx *gin.Context) {
	// Mendapatkan data bestseller dari request body
	var bestseller m_bestseller.Bestseller
	err := ctx.ShouldBind(&bestseller)
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

	// Memanggil service untuk update bestseller
	updatedBestseller, err := c.updateBestsellerService.UpdateBestseller(ctx, bestseller, file)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Menampilkan respon JSON
	ctx.JSON(http.StatusOK, gin.H{"data": updatedBestseller, "message": "Bestseller berhasil diperbarui"})
}
