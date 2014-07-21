package main

import (
	"github.com/dickeyxxx/updeps/languages"
	"github.com/dickeyxxx/updeps/pkg"
	"github.com/gin-gonic/gin"
	"labix.org/v2/mgo"
)

func main() {
	r := gin.Default()
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	db := session.DB("updeps")
	languagesClient := languages.NewClient(db.C("Languages"))
	pkgClient := pkg.NewClient(db.C("Packages"))
	r.Use(func(c *gin.Context) {
		c.Set("LanguageClient", languagesClient)
		c.Set("PkgClient", pkgClient)
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding")
	})
	Route(r)
	r.Run(":5001")
}
