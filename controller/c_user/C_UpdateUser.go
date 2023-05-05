package c_user

import (
	"net/http"
	"toko-buah/model/m_user"
	"toko-buah/service/s_user"

	"github.com/gin-gonic/gin"
)

type UpdateUserController struct {
	updateUserService *s_user.UpdateUserService
}

func NewUpdateUserController(updateUserService *s_user.UpdateUserService) *UpdateUserController {
	return &UpdateUserController{updateUserService}
}

func (c *UpdateUserController) UpdateUser(ctx *gin.Context) {
	// Mendapatkan data user dari request body
	var user m_user.User
	err := ctx.ShouldBind(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 	// Mendapatkan file foto profil dari request body
	file, err := ctx.FormFile("profile_picture")
	if err != nil && err != http.ErrMissingFile {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Memanggil service untuk update user
	updatedUser, err := c.updateUserService.UpdateUser(ctx, user, file)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Menampilkan respon JSON
	ctx.JSON(http.StatusOK, gin.H{"data": updatedUser, "message": "User berhasil diperbarui"})
}
