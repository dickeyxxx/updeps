package api

import (
	"github.com/dickeyxxx/updeps/category"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/binding"
)

func Initialize(m martini.Router) {
	m.Get("/packages", packageList)
	m.Get("/packages/**", packageDetail)
	m.Get("/categories", categoryList)
	m.Post("/categories", binding.Bind(category.Category{}), categoryCreate)

	m.Get("**", func() (int, string) {
		return 404, "Not Found"
	})
}
