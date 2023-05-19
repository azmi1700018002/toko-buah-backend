package c_home

import (
	"net/http"
	"strconv"
	"toko-buah/service/s_home"

	"github.com/gin-gonic/gin"
)

type deleteHomeController struct {
	deleteHomeService s_home.DeleteHomeService
}

func NewHomeDeleteController(deleteHomeService s_home.DeleteHomeService) *deleteHomeController {
	return &deleteHomeController{deleteHomeService}
}

func (c *deleteHomeController) DeleteHome(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("home_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	err = c.deleteHomeService.DeleteHomeByID(id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": "Home berhasil dihapus"})
}
