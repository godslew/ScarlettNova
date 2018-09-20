package main

import (
	"github.com/gin-gonic/gin"
	"github.com/godslew/ScarlettNova/controllers"
)

func main() {
	r := gin.Default()
	ping := controllers.NewPingController(r)
	ping.GetPing()
	ping.PostPing()

	twitter := controllers.NewTwitterController(r)
	twitter.GetCrcToken()
	r.Run(":19810")
}
