package category

import (
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

func (c *Client) All() ([]Category, error) {
	var result []Category
	err := c.collection().Find(bson.M{}).All(&result)
	return result, err
}

func (c *Client) UpsertCategory(category *Category) (*mgo.ChangeInfo, error) {
	return c.collection().Upsert(bson.M{"name": category.Name}, &category)
}

func (c *Client) collection() *mgo.Collection {
	return c.db.C("categories")
}
