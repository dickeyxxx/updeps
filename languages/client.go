package languages

import "labix.org/v2/mgo"

type ClientInterface interface {
	List() []Language
	Get(slug string) Language
}

type Client struct {
	db *mgo.Collection
}

func NewClient(db *mgo.Database) ClientInterface {
	return &Client{db.C("languages")}
}

func (c *Client) List() []Language {
	return []Language{
		Language{"Go", "go"},
		Language{"JavaScript", "javascript"},
		Language{"Ruby", "ruby"},
		Language{"Python", "python"},
	}
}

func (c *Client) Get(slug string) Language {
	return Language{"Go", "go"}
}
