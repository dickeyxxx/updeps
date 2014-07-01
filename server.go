package main

import (
	"github.com/dickeyxxx/updeps/api"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"labix.org/v2/mgo"
)

func main() {
	m := martini.Classic()
	m.Use(martini.Static("templates"))
	m.Use(martini.Static("assets"))
	m.Use(render.Renderer())

	mongo, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer mongo.Close()
	db := mongo.DB("updeps")
	m.Map(db)

	m.Group("/api", func(r martini.Router) {
		api.Initialize(r)
	})

	m.Get("**", func(r render.Render) {
		r.HTML(200, "app", nil)
	})
	m.Run()
}
