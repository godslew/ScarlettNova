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
	r.Run(":19810")
}
