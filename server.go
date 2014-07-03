package main

import (
	"github.com/dickeyxxx/updeps/api"
	"github.com/dickeyxxx/updeps/config"
	"github.com/dickeyxxx/updeps/models"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

func main() {
	m := martini.Classic()
	m.Use(martini.Static("templates"))
	m.Use(martini.Static("assets"))
	m.Use(render.Renderer())

	mongo, err := config.Mongo()
	if err != nil {
		panic(err)
	}
	defer mongo.Close()
	db := mongo.DB("updeps")
	models := models.NewClient(db)
	m.Map(models)

	m.Group("/api", func(r martini.Router) {
		api.Initialize(r)
	})

	m.Get("**", func(r render.Render) {
		r.HTML(200, "app", nil)
	})
	m.Run()
}
