package c_testimoni

import (
	"net/http"
	m_testimoni "toko-buah/model/m_testimoni"
	s_testimoni "toko-buah/service/s_testimoni"

	"github.com/gin-gonic/gin"
)

type addTestimoniController struct {
	addTestimoniService *s_testimoni.AddTestimoniService
}

func NewTestimoniAddController(addTestimoniService *s_testimoni.AddTestimoniService) *addTestimoniController {
	return &addTestimoniController{addTestimoniService}
}

func (c *addTestimoniController) AddTestimoni(ctx *gin.Context) {
	var testimoni m_testimoni.Testimoni
	if err := ctx.ShouldBind(&testimoni); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	addTestimoni, err := c.addTestimoniService.AddTestimoni(ctx, testimoni)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": addTestimoni, "message": "Testimoni Telah Berhasil Ditambahkan"})
}
