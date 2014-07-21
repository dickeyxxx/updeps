package main

import "github.com/gin-gonic/gin"

func Route(r *gin.Engine) {
	r.OPTIONS("*path", options)
	r.GET("/languages", languageList)
	r.GET("/languages/:language", languageShow)
	r.GET("/languages/:language/packages", pkgList)
	r.POST("/packages", pkgCreate)
}

func options(c *gin.Context) {
	c.String(200, "")
}
