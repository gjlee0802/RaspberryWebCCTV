package handler

import (
	"github.com/gin-gonic/gin"
)

func ShowIndexPage(c *gin.Context) {
	// Call the render function with the name of the template to render
	render(c, gin.H{"title":   "Home Page"}, "index.html");
}

func ShowStreamPage(c *gin.Context) {
	render(c, gin.H{"title":   "Stream Page"}, "stream.html");
}

func ShowVideoLogPage(c *gin.Context) {
	render(c, gin.H{"title":   "Video Log Page"}, "video_log.html");
}
