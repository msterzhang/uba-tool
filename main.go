package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)


func main() {
	r := gin.Default()
	r.Static("/statics", "./statics")
	r.LoadHTMLGlob("templates/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html",gin.H{"Host":"http://localhost:8006"})
	})
	r.Run(":8006")
	fmt.Println("http://localhost:8006")
}