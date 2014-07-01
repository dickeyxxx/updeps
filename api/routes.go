package api

import (
	"github.com/go-martini/martini"
	"github.com/google/go-github/github"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/render"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

func Initialize(m martini.Router) {
	m.Post("/packages", binding.Bind(Package{}), func(pkg Package, r render.Render, db *mgo.Database) {
		c := db.C("packages")

		client := github.NewClient(nil)
		repo, _, err := client.Repositories.Get(pkg.Owner, pkg.Name)
		if err != nil {
			panic(err)
		}

		pkg.Stars = *repo.StargazersCount
		pkg.Forks = *repo.ForksCount

		if _, err = c.Upsert(bson.M{"name": pkg.Name}, &pkg); err != nil {
			panic(err)
		}

		r.JSON(201, pkg)
	})

	m.Post("/packages/refresh", func(db *mgo.Database) int {
		Refresh(db)
		return 200
	})

	m.Get("/packages", func(r render.Render, db *mgo.Database) {
		c := db.C("packages")
		var result []Package
		c.Find(bson.M{}).Limit(1000).All(&result)
		r.JSON(200, result)
	})

	m.Get("/packages/**", func(r render.Render, params martini.Params, db *mgo.Database) {
		path := params["_1"]
		c := db.C("packages")
		var result Package
		c.Find(bson.M{"path": path}).One(&result)
		r.JSON(200, result)
	})
}
