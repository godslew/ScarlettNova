package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/godslew/ScarlettNova/parameters/responses"
)

type PingController struct {
	engine *gin.Engine
}

func NewPingController(engine *gin.Engine) *PingController {
	return &PingController{engine}
}

func (c *PingController) GetPing() {
	c.engine.GET("/ping", c.getPing)
}

func (c *PingController) getPing(ctx *gin.Context) {
	res := responses.GetPingResponse{Message: "pong", Test: "testetstetstet"}
	ctx.JSON(200, res)
}