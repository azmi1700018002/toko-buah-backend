package c_about

import (
	"net/http"
	"toko-buah/service/s_about"

	"github.com/gin-gonic/gin"
)

type AboutController struct {
	getAboutService s_about.GetAboutService
}

func NewGetAboutController(getAboutService s_about.GetAboutService) *AboutController {
	return &AboutController{getAboutService}
}

func (c *AboutController) GetAllAbout(ctx *gin.Context) {
	about, err := c.getAboutService.GetAllAbout()
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Hasil data about", "data": about})
}
