package c_newarrival

import (
	"net/http"
	"strconv"
	s_newarrival "toko-buah/service/s_new_arrival"

	"github.com/gin-gonic/gin"
)

type deleteNewArrivalController struct {
	deleteNewArrivalService s_newarrival.DeleteNewArrivalService
}

func NewNewArrivalDeleteController(deleteNewArrivalService s_newarrival.DeleteNewArrivalService) *deleteNewArrivalController {
	return &deleteNewArrivalController{deleteNewArrivalService}
}

func (c *deleteNewArrivalController) DeleteNewArrival(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("new_arrival_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	err = c.deleteNewArrivalService.DeleteNewArrivalByID(id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": "NewArrival berhasil dihapus"})
}
