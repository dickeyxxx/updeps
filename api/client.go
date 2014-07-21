package api

import (
	"github.com/dickeyxxx/updeps/languages"
	"github.com/dickeyxxx/updeps/pkg"
	"labix.org/v2/mgo"
)

type db interface {
	C(name string) *mgo.Collection
}

type Client struct {
	db        db
	pkg       *pkg.Client
	languages languages.ClientInterface
}

func NewClient(db db) *Client {
	return &Client{
		db:        db,
		pkg:       pkg.NewClient(db),
		languages: languages.NewClient(),
	}
}
