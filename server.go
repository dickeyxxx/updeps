package main

import (
	"github.com/dickeyxxx/updeps/api"
	"github.com/dickeyxxx/updeps/languages"
	"github.com/dickeyxxx/updeps/pkg"
	"github.com/gin-gonic/gin"
	"labix.org/v2/mgo"
)

type server struct {
	db       *mgo.Database
	language languages.ClientInterface
	pkg      *pkg.Client
}

func main() {
	r := gin.Default()
	api := api.NewClient(db())
	api.Route(r)
	r.Run(":5001")
}

func db() *mgo.Database {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	return session.DB("updeps")
}
