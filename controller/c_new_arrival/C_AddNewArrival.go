package c_newarrival

import (
	"net/http"
	m_newarrival "toko-buah/model/m_new_arrival"
	s_newarrival "toko-buah/service/s_new_arrival"

	"github.com/gin-gonic/gin"
)

type addNewArrivalController struct {
	addNewArrivalService *s_newarrival.AddNewArrivalService
}

func NewNewArrivalAddController(addNewArrivalService *s_newarrival.AddNewArrivalService) *addNewArrivalController {
	return &addNewArrivalController{addNewArrivalService}
}

func (c *addNewArrivalController) AddNewArrival(ctx *gin.Context) {
	var newarrival m_newarrival.NewArrival
	if err := ctx.ShouldBind(&newarrival); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	addNewArrival, err := c.addNewArrivalService.AddNewArrival(ctx, newarrival)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": addNewArrival, "message": "NewArrival Telah Berhasil Ditambahkan"})
}
