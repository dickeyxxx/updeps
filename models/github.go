package models

import (
	"fmt"
	"os"
	"time"

	"code.google.com/p/goauth2/oauth"

	"github.com/google/go-github/github"
)

func (c *Client) RefreshPackagesGithub() {
	packageChannel := make(chan Package)
	t := &oauth.Transport{
		Token: &oauth.Token{AccessToken: os.Getenv("GITHUB_ACCESS_TOKEN")},
	}
	githubClient := github.NewClient(t.Client())
	go c.githubWorker(githubClient, packageChannel)
	go c.githubWorker(githubClient, packageChannel)
	go c.githubWorker(githubClient, packageChannel)
	packages, err := c.AllPackages()
	if err != nil {
		panic(err)
	}
	for _, pkg := range packages {
		if pkg.GithubName != "" && pkg.GithubUpdatedAt.IsZero() {
			packageChannel <- pkg
		}
	}
}

func (c *Client) githubWorker(githubClient *github.Client, packages <-chan Package) {
	for pkg := range packages {
		repo, resp, err := githubClient.Repositories.Get(pkg.GithubOwner, pkg.GithubName)
		if err != nil {
			githubError, ok := err.(*github.ErrorResponse)
			if ok && githubError.Response.StatusCode == 403 {
				fmt.Println("Rate limited...")
				resp.Body.Close()
				time.Sleep(10 * time.Second)
				continue
			} else {
				panic(err)
			}
		}
		pkg.GithubForks = *repo.ForksCount
		pkg.GithubStargazers = *repo.StargazersCount
		pkg.GithubUpdatedAt = time.Now()
		fmt.Println("updating package", pkg)
		if _, err := c.UpsertPackage(&pkg); err != nil {
			panic(err)
		}
	}
}
