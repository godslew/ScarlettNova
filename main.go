package main

import (
	"github.com/gin-gonic/gin"
	"github.com/godslew/ScarlettNova/controllers"
)

func main() {
	r := gin.Default()
	twitter := controllers.NewTwitterController(r)
	twitter.GetCrcToken()
	twitter.PostWebHook()
	twitter.PostWebHookTest()
	r.Run(":19810")
}
