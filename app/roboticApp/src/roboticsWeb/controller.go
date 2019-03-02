package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func registerRoutes() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("templates/**/**.html")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.Static("/public", "./public")
	
	return r
}
