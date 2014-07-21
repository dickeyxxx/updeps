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
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	api := api.NewClient(session.DB("updeps"))
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding")
	})
	api.Route(r)
	r.Run(":5001")
}
