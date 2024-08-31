package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ntquang/ecommerce/internal/services"
)

type PongController struct {
	pongService *services.PongService
}

func NewPongController() *PongController {
	return &PongController{
		pongService: services.NewPongService(),
	}
}

func (p *PongController) Pong(c *gin.Context) {
	fmt.Println("router")
	age := c.Query("age")
	c.JSON(http.StatusOK, gin.H{
		"message": "pong " + p.pongService.GetPong(),
		"age":     "is " + age,
	})
}
