package pkg

import (
	"fmt"
	"time"
)

func (c *Client) RefreshPackagesGithub() {
	packageChannel := make(chan Package)
	for i := 0; i < 10; i++ {
		go c.githubWorker(packageChannel)
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

func (c *Client) githubWorker(packages <-chan Package) {
	for pkg := range packages {
		repo, err := c.github.GetRepoInfo(pkg.GithubOwner, pkg.GithubName)
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
