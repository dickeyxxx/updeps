package category

import "labix.org/v2/mgo"

type Client struct {
	db *mgo.Database
}

func NewClient(db *mgo.Database) *Client {
	c := &Client{db}
	return c
}
