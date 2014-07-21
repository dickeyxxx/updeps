package api

import (
	"github.com/dickeyxxx/updeps/languages"
	"github.com/dickeyxxx/updeps/pkg"
	"labix.org/v2/mgo"
)

type Client struct {
	db        *mgo.Database
	pkg       *pkg.Client
	languages languages.ClientInterface
}

func NewClient(db *mgo.Database) *Client {
	return &Client{
		db:        db,
		pkg:       pkg.NewClient(db),
		languages: languages.NewClient(db),
	}
}
