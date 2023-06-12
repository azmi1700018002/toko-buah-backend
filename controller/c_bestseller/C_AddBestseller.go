package c_bestseller

import (
	"net/http"
	m_bestseller "toko-buah/model/m_best_seller"
	s_bestseller "toko-buah/service/s_best_seller"

	"github.com/gin-gonic/gin"
)

type addBestsellerController struct {
	addBestsellerService *s_bestseller.AddBestsellerService
}

func NewBestsellerAddController(addBestsellerService *s_bestseller.AddBestsellerService) *addBestsellerController {
	return &addBestsellerController{addBestsellerService}
}

func (c *addBestsellerController) AddBestseller(ctx *gin.Context) {
	var bestseller m_bestseller.Bestseller
	if err := ctx.ShouldBind(&bestseller); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	addBestseller, err := c.addBestsellerService.AddBestseller(ctx, bestseller)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": addBestseller, "message": "Bestseller Telah Berhasil Ditambahkan"})
}
