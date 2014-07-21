package api

import (
	"github.com/dickeyxxx/updeps/pkg"
	"github.com/gin-gonic/gin"
)

func (s *Client) pkgCreate(c *gin.Context) {
	var pkg pkg.Pkg

	if c.EnsureBody(&pkg) {
		s.pkg.Create(&pkg)
		c.JSON(201, pkg)
	}
}

func (s *Client) pkgList(c *gin.Context) {
	packages, err := s.pkg.List()
	if err != nil {
		panic(err)
	}
	c.JSON(200, packages)
}
