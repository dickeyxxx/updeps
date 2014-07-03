package pkg

import (
	"time"

	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

type Package struct {
	Path             string    `bson:"path,omitempty"`
	GithubName       string    `bson:"github_name,omitempty"`
	GithubOwner      string    `bson:"github_owner,omitempty"`
	GithubUpdatedAt  time.Time `bson:"github_updated_at,omitempty"`
	GithubForks      int       `bson:"github_forks,omitempty"`
	GithubStargazers int       `bson:"github_stargazers,omitempty"`
}

func (c *Client) AllPackages() ([]Package, error) {
	var result []Package
	err := c.packagesCollection().Find(bson.M{}).All(&result)
	return result, err
}

func (c *Client) PackagesByStars(limit int) ([]Package, error) {
	var result []Package
	err := c.packagesCollection().Find(bson.M{}).Sort("-github_stargazers").Limit(limit).All(&result)
	return result, err
}

func (c *Client) PackageByPath(path string) (*Package, error) {
	var result Package
	err := c.packagesCollection().Find(bson.M{"path": path}).One(&result)
	return &result, err
}

func (c *Client) UpsertPackage(pkg *Package) (*mgo.ChangeInfo, error) {
	return c.packagesCollection().Upsert(bson.M{"path": pkg.Path}, &pkg)
}

func (c *Client) packagesCollection() *mgo.Collection {
	return c.db.C("packages")
}
