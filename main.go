package main

import (
	"github.com/gin-gonic/gin"
	"strings"
)

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{})
	})

	r.GET("/tolower", func(c *gin.Context) {
		c.HTML(200, "tolower.html", gin.H{})
	})

	r.POST("/tolower", func(c *gin.Context) {
		info := c.PostForm("info")
		result := strings.ToLower(info)
		c.HTML(200, "tolower.html", gin.H{"result": result})
	})

	r.GET("/toupper", func(c *gin.Context) {
		c.HTML(200, "toupper.html", gin.H{})
	})

	r.POST("/toupper", func(c *gin.Context) {
		info := c.PostForm("info")
		result := strings.ToUpper(info)
		c.HTML(200, "toupper.html", gin.H{"result": result})
	})

	r.Run(":8080")
}
