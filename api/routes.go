package api

import (
	"github.com/dickeyxxx/updeps/models"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/render"
)

func Initialize(m martini.Router) {
	m.Post("/packages", binding.Bind(models.Package{}), func(pkg models.Package, r render.Render, models *models.Client) {
		if _, err := models.UpsertPackage(&pkg); err != nil {
			panic(err)
		}
		r.JSON(201, pkg)
	})

	m.Get("/packages", func(r render.Render, models *models.Client) {
		packages, err := models.PackagesByStars()
		if err != nil {
			panic(err)
		}
		r.JSON(200, packages)
	})

	m.Get("/packages/**", func(r render.Render, params martini.Params, models *models.Client) {
		path := params["_1"]
		result, err := models.PackageByPath(path)
		if err != nil {
			panic(err)
		}
		r.JSON(200, result)
	})

	m.Get("**", func() (int, string) {
		return 404, "Not Found"
	})
}
