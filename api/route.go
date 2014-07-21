package api

import "github.com/gin-gonic/gin"

type ginEngineI interface {
	GET(path string, handlers ...gin.HandlerFunc)
	POST(path string, handlers ...gin.HandlerFunc)
	OPTIONS(path string, handlers ...gin.HandlerFunc)
	Use(middlewars ...gin.HandlerFunc)
}

func (s *Client) Route(r ginEngineI) {
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding")
	})
	r.GET("/languages", s.languageList)
	r.GET("/languages/:language", s.languageShow)
	r.GET("/languages/:language/packages", s.pkgList)
	r.POST("/packages", s.pkgCreate)
	r.OPTIONS("*path", s.options)
}

func (s *Client) options(c *gin.Context) {
	c.String(200, "")
}
