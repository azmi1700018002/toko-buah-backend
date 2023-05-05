package c_buah

import (
	"net/http"
	"toko-buah/model/m_buah"
	"toko-buah/service/s_buah"

	"github.com/gin-gonic/gin"
)

type UpdateBuahController struct {
	updateBuahService *s_buah.UpdateBuahService
}

func NewUpdateBuahController(updateBuahService *s_buah.UpdateBuahService) *UpdateBuahController {
	return &UpdateBuahController{updateBuahService}
}

func (c *UpdateBuahController) UpdateBuah(ctx *gin.Context) {
	// Mendapatkan data buah dari request body
	var buah m_buah.Buah
	err := ctx.ShouldBind(&buah)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Memanggil service untuk update buah
	updatedBuah, err := c.updateBuahService.UpdateBuah(ctx, buah)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Menampilkan respon JSON
	ctx.JSON(http.StatusOK, gin.H{"data": updatedBuah, "message": "Buah berhasil diupdate"})
}
