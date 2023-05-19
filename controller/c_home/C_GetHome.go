package c_home

import (
	"net/http"
	"toko-buah/service/s_home"

	"github.com/gin-gonic/gin"
)

type HomeController struct {
	getHomeService s_home.GetHomeService
}

func NewGetHomeController(getHomeService s_home.GetHomeService) *HomeController {
	return &HomeController{getHomeService}
}

func (c *HomeController) GetAllHome(ctx *gin.Context) {
	home, err := c.getHomeService.GetAllHome()
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Hasil data home", "data": home})
}
