package main

import (
	"github.com/dickeyxxx/updeps/pkg"
	"github.com/gin-gonic/gin"
)

func pkgCreate(c *gin.Context) {
	client := c.MustGet("PkgClient").(*pkg.Client)
	var pkg pkg.Pkg

	if c.EnsureBody(&pkg) {
		client.Create(&pkg)
		c.JSON(201, pkg)
	}
}

func pkgList(c *gin.Context) {
	client := c.MustGet("PkgClient").(*pkg.Client)
	packages, err := client.List()
	if err != nil {
		panic(err)
	}
	c.JSON(200, packages)
}
