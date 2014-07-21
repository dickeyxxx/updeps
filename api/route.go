package api

import "github.com/gin-gonic/gin"

type ginEngineI interface {
	GET(path string, handlers ...gin.HandlerFunc)
	POST(path string, handlers ...gin.HandlerFunc)
	OPTIONS(path string, handlers ...gin.HandlerFunc)
}

func (s *Client) Route(r ginEngineI) {
	r.GET("/languages", s.languageList)
	r.GET("/languages/:language", s.languageShow)
	r.GET("/languages/:language/packages", s.pkgList)
	r.POST("/packages", s.pkgCreate)
	r.OPTIONS("*path", s.options)
}

func (s *Client) options(c *gin.Context) {
	c.String(200, "")
}
