package pkg

import (
	"github.com/dickeyxxx/updeps/github"
	"labix.org/v2/mgo"
)

type Client struct {
	db     *mgo.Database
	github *github.Client
}

func NewClient(db *mgo.Database, github *github.Client) *Client {
	c := &Client{db, github}
	return c
}
