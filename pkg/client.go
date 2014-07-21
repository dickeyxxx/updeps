package pkg

import (
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

type Client struct {
	db *mgo.Collection
}

func NewClient(db *mgo.Database) *Client {
	return &Client{db.C("packages")}
}

func (c *Client) Create(p *Pkg) error {
	return c.db.Insert(p)
}

func (c *Client) List() ([]Pkg, error) {
	var result []Pkg
	err := c.db.Find(bson.M{}).All(&result)
	return result, err
}
