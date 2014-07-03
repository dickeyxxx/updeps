package pkg

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
)

type response struct {
	Results []Package
}

func (c *Client) RefreshPackages() {
	packages := make(chan Package)
	for i := 1; i <= 10; i++ {
		go c.godocWorker(packages)
	}
	fetchGodocPackages(packages)
}

func (c *Client) godocWorker(packages <-chan Package) {
	for pkg := range packages {
		fmt.Println("adding package", pkg)
		if _, err := c.UpsertPackage(&pkg); err != nil {
			panic(err)
		}
	}
}

func fetchGodocPackages(packageChannel chan Package) {
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
	r := regexp.MustCompile("github.com/([^/]+)/([^/]+)")
	for _, pkg := range packages.Results {
		githubPath := r.FindStringSubmatch(pkg.Path)

		if len(githubPath) != 0 {
			pkg.GithubOwner = githubPath[1]
			pkg.GithubName = githubPath[2]
		}

		packageChannel <- pkg
	}
}
