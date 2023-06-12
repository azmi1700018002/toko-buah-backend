package c_testimoni

import (
	"net/http"
	m_testimoni "toko-buah/model/m_testimoni"
	s_testimoni "toko-buah/service/s_testimoni"

	"github.com/gin-gonic/gin"
)

type UpdateTestimoniController struct {
	updateTestimoniService *s_testimoni.UpdateTestimoniService
}

func NewUpdateTestimoniController(updateTestimoniService *s_testimoni.UpdateTestimoniService) *UpdateTestimoniController {
	return &UpdateTestimoniController{updateTestimoniService}
}

func (c *UpdateTestimoniController) UpdateTestimoni(ctx *gin.Context) {
	// Mendapatkan data testimoni dari request body
	var testimoni m_testimoni.Testimoni
	err := ctx.ShouldBind(&testimoni)
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

	// Memanggil service untuk update testimoni
	updatedTestimoni, err := c.updateTestimoniService.UpdateTestimoni(ctx, testimoni, file)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Menampilkan respon JSON
	ctx.JSON(http.StatusOK, gin.H{"data": updatedTestimoni, "message": "Testimoni berhasil diupdate"})
}
