package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"

	"github.com/dickeyxxx/updeps/models"
)

type response struct {
	Results []models.Package
}

func Refresh(c *models.Client) {
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
	for _, pkg := range packages.Results {
		githubPath := r.FindStringSubmatch(pkg.Path)

		if len(githubPath) != 0 {
			pkg.GithubOwner = githubPath[1]
			pkg.GithubName = githubPath[2]
		}
		fmt.Println("Adding", pkg)
		if _, err = c.UpsertPackage(&pkg); err != nil {
			panic(err)
		}
	}
}
