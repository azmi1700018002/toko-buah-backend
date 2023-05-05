package c_user

import (
	"net/http"
	"toko-buah/service/s_user"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type deleteUserController struct {
	deleteUserService s_user.DeleteUserService
}

func NewUserDeleteController(deleteUserService s_user.DeleteUserService) *deleteUserController {
	return &deleteUserController{deleteUserService}
}

func (c *deleteUserController) DeleteUser(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id_user"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	err = c.deleteUserService.DeleteUserByID(id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": "User berhasil dihapus"})
}
