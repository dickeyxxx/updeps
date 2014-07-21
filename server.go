package main

import (
	"log"
	"net/http"

	"github.com/dickeyxxx/updeps/api"
	"github.com/dickeyxxx/updeps/category"
	"github.com/dickeyxxx/updeps/config"
	"github.com/dickeyxxx/updeps/pkg"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

func main() {
	m := martini.Classic()
	m.Use(render.Renderer())
	m.Use(func(res http.ResponseWriter) {
		res.Header().Set("Access-Control-Allow-Origin", "*")
	})

	mongo, err := config.Mongo()
	if err != nil {
		panic(err)
	}
	defer mongo.Close()
	db := mongo.DB("updeps")
	pkg := pkg.NewClient(db, nil)
	category := category.NewClient(db)
	m.Map(pkg)
	m.Map(category)
	api.Initialize(m)

	log.Fatal(http.ListenAndServe(":5001", m))
}
