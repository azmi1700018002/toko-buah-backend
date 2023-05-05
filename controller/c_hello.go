package controller

import "github.com/gin-gonic/gin"

// Make swagger documentation for this function

// @Summary Hello World
// @Description Hello World
// @Tags Hello
// @Accept  json
// @Produce  json
// @Router / [get]
func Helloworld(g *gin.Context) {
	g.JSON(200, gin.H{
		"message": "It works",
	})
}
