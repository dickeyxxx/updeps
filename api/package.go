package api

import (
	"github.com/dickeyxxx/updeps/pkg"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

func packageList(r render.Render, pkg *pkg.Client) {
	result, err := pkg.PackagesByStars(10)
	if err != nil {
		panic(err)
	}
	r.JSON(200, result)
}

func packageDetail(r render.Render, params martini.Params, pkg *pkg.Client) {
	path := params["_1"]
	result, err := pkg.PackageByPath(path)
	if err != nil {
		panic(err)
	}
	r.JSON(200, result)
}
