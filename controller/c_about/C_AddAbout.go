package c_about

import (
	"net/http"
	"toko-buah/model/m_about"
	"toko-buah/service/s_about"

	"github.com/gin-gonic/gin"
)

type addAboutController struct {
	addAboutService *s_about.AddAboutService
}

func NewAboutAddController(addAboutService *s_about.AddAboutService) *addAboutController {
	return &addAboutController{addAboutService}
}

func (c *addAboutController) AddAbout(ctx *gin.Context) {
	var about m_about.About
	if err := ctx.ShouldBind(&about); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	addAbout, err := c.addAboutService.AddAbout(ctx, about)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": addAbout, "message": "About Telah Berhasil Ditambahkan"})
}
