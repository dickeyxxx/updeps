package api

import (
	"github.com/dickeyxxx/updeps/pkg"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/render"
)

func Initialize(m martini.Router) {
	m.Post("/packages", binding.Bind(pkg.Package{}), func(pkg pkg.Package, r render.Render, pkgClient *pkg.Client) {
		if _, err := pkgClient.UpsertPackage(&pkg); err != nil {
			panic(err)
		}
		r.JSON(201, pkg)
	})

	m.Get("/packages", func(r render.Render, pkg *pkg.Client) {
		packages, err := pkg.PackagesByStars()
		if err != nil {
			panic(err)
		}
		r.JSON(200, packages)
	})

	m.Get("/packages/**", func(r render.Render, params martini.Params, pkg *pkg.Client) {
		path := params["_1"]
		result, err := pkg.PackageByPath(path)
		if err != nil {
			panic(err)
		}
		r.JSON(200, result)
	})

	m.Get("**", func() (int, string) {
		return 404, "Not Found"
	})
}
