package api

import "github.com/gin-gonic/gin"

func (s *Client) languageList(c *gin.Context) {
	c.JSON(200, s.languages.List())
}

func (s *Client) languageShow(c *gin.Context) {
	c.JSON(200, s.languages.Get(c.Params.ByName("language")))
}
