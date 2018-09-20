package controllers

import (
	"fmt"
	"net/http"

	"github.com/godslew/ScarlettNova/parameters/requests"

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
	req := requests.GetPingRequest{}
	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}
	fmt.Println(req)
	res := responses.GetPingResponse{Message: "pong", Test: req.Message}
	ctx.JSON(http.StatusOK, res)
}

func (c *PingController) PostPing() {
	c.engine.POST("/ping", c.postPing)
}

func (c *PingController) postPing(ctx *gin.Context) {
	req := requests.PostPingRequest{}
	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}
	fmt.Println(req)
	res := responses.GetPingResponse{Message: "pong", Test: req.Message}
	ctx.JSON(http.StatusOK, res)
}
