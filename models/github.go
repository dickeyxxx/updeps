package models

import (
	"fmt"
	"time"

	"github.com/dickeyxxx/updeps/config"
	"github.com/dickeyxxx/updeps/github"
)

func (c *Client) RefreshPackagesGithub() {
	packageChannel := make(chan Package)
	githubClient := config.Github()
	for i := 0; i < 10; i++ {
		go c.githubWorker(githubClient, packageChannel)
	}
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

func (c *Client) githubWorker(github *github.Client, packages <-chan Package) {
	for pkg := range packages {
		repo, err := github.GetRepoInfo(pkg.GithubOwner, pkg.GithubName)
		if err != nil {
			panic(err)
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
