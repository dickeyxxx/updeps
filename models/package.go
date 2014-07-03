package models

import (
	"time"

	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

type Package struct {
	Path             string
	GithubName       string
	GithubOwner      string
	GithubUpdatedAt  time.Time
	GithubForks      int
	GithubStargazers int
}

func (c *Client) AllPackages() ([]Package, error) {
	var result []Package
	err := c.packagesCollection().Find(bson.M{}).All(&result)
	return result, err
}

func (c *Client) PackagesByStars() ([]Package, error) {
	var result []Package
	err := c.packagesCollection().Find(bson.M{}).Sort("-githubstargazers").Limit(1000).All(&result)
	return result, err
}

func (c *Client) PackageByPath(path string) (*Package, error) {
	var result Package
	err := c.packagesCollection().Find(bson.M{"path": path}).One(&result)
	return &result, err
}

func (c *Client) UpsertPackage(pkg *Package) (*mgo.ChangeInfo, error) {
	return c.packagesCollection().Upsert(bson.M{"Path": pkg.Path}, &pkg)
}

func (c *Client) packagesCollection() *mgo.Collection {
	return c.db.C("packages")
}
