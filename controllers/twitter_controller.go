package controllers

import (
	"fmt"
	"net/http"

	"github.com/godslew/ScarlettNova/libs/twitter"

	"github.com/gin-gonic/gin"
	"github.com/godslew/ScarlettNova/parameters/requests"
	"github.com/godslew/ScarlettNova/parameters/responses"
)

type TwitterController struct {
	engine *gin.Engine
}

func NewTwitterController(engine *gin.Engine) *TwitterController {
	return &TwitterController{engine}
}

func (c *TwitterController) GetCrcToken() {
	c.engine.GET("/twitter_webhook", c.getCrcToken)
}

func (c *TwitterController) getCrcToken(ctx *gin.Context) {
	req := requests.NewGetTwitterWebhookRequest()
	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}

	fmt.Print(req)

	res := responses.NewGetTwitterWebHookCrcCheckResponse()
	res.Token = twitter.CreateCRCToken(req.CrcToken)
	ctx.JSON(http.StatusOK, res)
}
