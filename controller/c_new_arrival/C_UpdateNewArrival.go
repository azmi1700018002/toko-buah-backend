package c_newarrival

import (
	"net/http"
	m_newarrival "toko-buah/model/m_new_arrival"
	s_newarrival "toko-buah/service/s_new_arrival"

	"github.com/gin-gonic/gin"
)

type UpdateNewArrivalController struct {
	updateNewArrivalService *s_newarrival.UpdateNewArrivalService
}

func NewUpdateNewArrivalController(updateNewArrivalService *s_newarrival.UpdateNewArrivalService) *UpdateNewArrivalController {
	return &UpdateNewArrivalController{updateNewArrivalService}
}

func (c *UpdateNewArrivalController) UpdateNewArrival(ctx *gin.Context) {
	// Mendapatkan data newarrival dari request body
	var newarrival m_newarrival.NewArrival
	err := ctx.ShouldBind(&newarrival)
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

	// Memanggil service untuk update newarrival
	updatedNewArrival, err := c.updateNewArrivalService.UpdateNewArrival(ctx, newarrival, file)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Menampilkan respon JSON
	ctx.JSON(http.StatusOK, gin.H{"data": updatedNewArrival, "message": "NewArrival berhasil diupdate"})
}
