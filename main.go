package main

import (
	"github.com/gjlee0802/RaspberryWebCCTV/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.LoadHTMLGlob("view/*")

	r.GET("/", handler.ShowIndexPage)
	r.GET("/index.html", handler.ShowIndexPage)
	r.GET("/stream.html", handler.ShowStreamPage)
	r.GET("/video_log.html", handler.ShowVideoLogPage)
	r.Run(":8080")
}
