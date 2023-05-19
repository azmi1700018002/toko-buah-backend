package c_home

import (
	"net/http"
	"toko-buah/model/m_home"
	"toko-buah/service/s_home"

	"github.com/gin-gonic/gin"
)

type UpdateHomeController struct {
	updateHomeService *s_home.UpdateHomeService
}

func NewUpdateHomeController(updateHomeService *s_home.UpdateHomeService) *UpdateHomeController {
	return &UpdateHomeController{updateHomeService}
}

func (c *UpdateHomeController) UpdateHome(ctx *gin.Context) {
	// Mendapatkan data home dari request body
	var home m_home.Home
	err := ctx.ShouldBind(&home)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Memanggil service untuk update home
	updatedHome, err := c.updateHomeService.UpdateHome(ctx, home)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Menampilkan respon JSON
	ctx.JSON(http.StatusOK, gin.H{"data": updatedHome, "message": "Home berhasil diupdate"})
}
