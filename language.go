package main

import (
	"github.com/dickeyxxx/updeps/languages"
	"github.com/gin-gonic/gin"
)

func languageList(c *gin.Context) {
	client := c.MustGet("LanguageClient").(languages.ClientInterface)
	c.JSON(200, client.List())
}

func languageShow(c *gin.Context) {
	client := c.MustGet("LanguageClient").(languages.ClientInterface)
	c.JSON(200, client.Get(c.Params.ByName("language")))
}
