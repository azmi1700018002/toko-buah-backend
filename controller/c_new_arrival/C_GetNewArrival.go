package c_newarrival

import (
	"net/http"
	s_newarrival "toko-buah/service/s_new_arrival"

	"github.com/gin-gonic/gin"
)

type NewArrivalController struct {
	getNewArrivalService s_newarrival.GetNewArrivalService
}

func NewGetNewArrivalController(getNewArrivalService s_newarrival.GetNewArrivalService) *NewArrivalController {
	return &NewArrivalController{getNewArrivalService}
}

func (c *NewArrivalController) GetAllNewArrival(ctx *gin.Context) {
	newarrival, err := c.getNewArrivalService.GetAllNewArrival()
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	count := len(newarrival) // Calculate the count of the products

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Hasil data new arrival",
		"data":    newarrival,
		"count":   count, // Add the count to the response
	})
}
