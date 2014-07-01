package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"

	"github.com/google/go-github/github"

	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

type response struct {
	Results []Package
}

func Refresh(db *mgo.Database) {
	client := &http.Client{}
	resp, err := client.Get("http://api.godoc.org/packages")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	var packages response
	if err := decoder.Decode(&packages); err != nil {
		panic(err)
	}
	r := regexp.MustCompile("github.com/(.*)/(.*)")
	c := db.C("packages")
	githubClient := github.NewClient(nil)
	for _, pkg := range packages.Results {
		githubPath := r.FindStringSubmatch(pkg.Path)

		if len(githubPath) != 0 {
			githubInfo := fetchGithubInfo(githubPath[1], githubPath[2], githubClient)
			pkg.Owner = *githubInfo.Owner.Name
			pkg.Name = *githubInfo.Name
			pkg.Forks = *githubInfo.ForksCount
			pkg.Stars = *githubInfo.StargazersCount
		}
		fmt.Println("Adding", pkg)
		if _, err = c.Upsert(bson.M{"Path": pkg.Path}, &pkg); err != nil {
			panic(err)
		}
	}
}

func fetchGithubInfo(owner string, name string, client *github.Client) *github.Repository {
	repo, _, err := client.Repositories.Get(owner, name)
	if err != nil {
		panic(err)
	}
	return repo
}
