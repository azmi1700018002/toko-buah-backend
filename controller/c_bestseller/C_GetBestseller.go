package c_bestseller

import (
	"net/http"
	s_bestseller "toko-buah/service/s_best_seller"

	"github.com/gin-gonic/gin"
)

type BestsellerController struct {
	getBestsellerService s_bestseller.GetBestsellerService
}

func NewGetBestsellerController(getBestsellerService s_bestseller.GetBestsellerService) *BestsellerController {
	return &BestsellerController{getBestsellerService}
}

func (c *BestsellerController) GetAllBestseller(ctx *gin.Context) {
	bestseller, err := c.getBestsellerService.GetAllBestseller()
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	count := len(bestseller) // Calculate the count of the products

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Hasil data bestseller",
		"data":    bestseller,
		"count":   count, // Add the count to the response
	})
}
