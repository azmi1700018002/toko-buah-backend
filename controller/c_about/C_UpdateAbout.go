package c_about

import (
	"net/http"
	"toko-buah/model/m_about"
	"toko-buah/service/s_about"

	"github.com/gin-gonic/gin"
)

type UpdateAboutController struct {
	updateAboutService *s_about.UpdateAboutService
}

func NewUpdateAboutController(updateAboutService *s_about.UpdateAboutService) *UpdateAboutController {
	return &UpdateAboutController{updateAboutService}
}

func (c *UpdateAboutController) UpdateAbout(ctx *gin.Context) {
	// Mendapatkan data about dari request body
	var about m_about.About
	err := ctx.ShouldBind(&about)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Memanggil service untuk update about
	updatedAbout, err := c.updateAboutService.UpdateAbout(ctx, about)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Menampilkan respon JSON
	ctx.JSON(http.StatusOK, gin.H{"data": updatedAbout, "message": "About berhasil diupdate"})
}
