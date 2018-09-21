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

func (c *TwitterController) PostWebHook() {
	c.engine.POST("/twitter_webhook", c.postWebHook)
}

func (c *TwitterController) postWebHook(ctx *gin.Context) {
	req := requests.NewPostTwitterWebhookRequest()
	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}

	fmt.Print(req)
	client, err := twitter.CreateTwitterClient()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}
	for _, e := range req.DirectMessageEvents {
		if e.Type != "message_create" {
			continue
		}
		text := "仕事お疲れ様！頑張ったね！！"
		dmReq := requests.NewPostDirectMessageTestRequest(e.MessageCreate.SenderID, text)
		fmt.Print(dmReq)
		if err := twitter.PostDM(client, dmReq); err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
		}
	}
	ctx.JSON(http.StatusOK, "")
}

func (c *TwitterController) PostWebHookTest() {
	c.engine.POST("/twitter_webhook/test", c.postWebHookTest)
}

func (c *TwitterController) postWebHookTest(ctx *gin.Context) {
	req := requests.NewPostTwitterWebHookTestRequest()
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}

	fmt.Print(req)

	client, err := twitter.CreateTwitterClient()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	dmReq := requests.NewPostDirectMessageTestRequest(req.ID, req.Message)
	fmt.Print(dmReq)
	if err := twitter.PostDM(client, dmReq); err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}
	ctx.JSON(http.StatusNoContent, "")
}
