package c_home

import (
	"net/http"
	"toko-buah/model/m_home"
	"toko-buah/service/s_home"

	"github.com/gin-gonic/gin"
)

type addHomeController struct {
	addHomeService *s_home.AddHomeService
}

func NewHomeAddController(addHomeService *s_home.AddHomeService) *addHomeController {
	return &addHomeController{addHomeService}
}

func (c *addHomeController) AddHome(ctx *gin.Context) {
	var home m_home.Home
	if err := ctx.ShouldBind(&home); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	addHome, err := c.addHomeService.AddHome(ctx, home)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": addHome, "message": "Home Telah Berhasil Ditambahkan"})
}
