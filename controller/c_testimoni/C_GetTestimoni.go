package c_testimoni

import (
	"net/http"
	s_testimoni "toko-buah/service/s_testimoni"

	"github.com/gin-gonic/gin"
)

type TestimoniController struct {
	getTestimoniService s_testimoni.GetTestimoniService
}

func NewGetTestimoniController(getTestimoniService s_testimoni.GetTestimoniService) *TestimoniController {
	return &TestimoniController{getTestimoniService}
}

func (c *TestimoniController) GetAllTestimoni(ctx *gin.Context) {
	testimoni, err := c.getTestimoniService.GetAllTestimoni()
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	count := len(testimoni) // Calculate the count of the products

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Hasil data testimoni",
		"data":    testimoni,
		"count":   count, // Add the count to the response
	})
}
