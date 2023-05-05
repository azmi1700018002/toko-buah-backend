package c_buah

import (
	"net/http"
	"toko-buah/service/s_buah"

	"github.com/gin-gonic/gin"
)

type BuahController struct {
	getBuahService s_buah.GetBuahService
}

func NewGetBuahController(getBuahService s_buah.GetBuahService) *BuahController {
	return &BuahController{getBuahService}
}

func (c *BuahController) GetAllBuah(ctx *gin.Context) {
	buah, err := c.getBuahService.GetAllBuah()
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Hasil data buah", "data": buah})
}
