package c_about

import (
	"net/http"
	"strconv"
	"toko-buah/service/s_about"

	"github.com/gin-gonic/gin"
)

type deleteAboutController struct {
	deleteAboutService s_about.DeleteAboutService
}

func NewAboutDeleteController(deleteAboutService s_about.DeleteAboutService) *deleteAboutController {
	return &deleteAboutController{deleteAboutService}
}

func (c *deleteAboutController) DeleteAbout(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("about_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	err = c.deleteAboutService.DeleteAboutByID(id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": "About berhasil dihapus"})
}
