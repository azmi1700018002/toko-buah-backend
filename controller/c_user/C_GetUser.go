package c_user

import (
	"net/http"
	"toko-buah/service/s_user"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserController struct {
	getUserService s_user.GetUserService
}

func NewGetUserController(getUserService s_user.GetUserService) *UserController {
	return &UserController{getUserService}
}

func (c *UserController) GetAllUser(ctx *gin.Context) {
	users, err := c.getUserService.GetAllUser()
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	ctx.JSON(http.StatusOK, users)
}

func (c *UserController) GetUserByID(ctx *gin.Context) {
	userID, err := uuid.Parse(ctx.Param("id_user"))
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	user, err := c.getUserService.GetUserByID(userID)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	ctx.JSON(http.StatusOK, user)
}
