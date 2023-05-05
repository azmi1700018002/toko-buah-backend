package c_buah

import (
	"net/http"
	"toko-buah/model/m_buah"
	"toko-buah/service/s_buah"

	"github.com/gin-gonic/gin"
)

type addBuahController struct {
	addBuahService *s_buah.AddBuahService
}

func NewBuahAddController(addBuahService *s_buah.AddBuahService) *addBuahController {
	return &addBuahController{addBuahService}
}

func (c *addBuahController) AddBuah(ctx *gin.Context) {
	var buah m_buah.Buah
	if err := ctx.ShouldBind(&buah); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	addBuah, err := c.addBuahService.AddBuah(ctx, buah)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": addBuah, "message": "Buah Telah Berhasil Ditambahkan"})
}
