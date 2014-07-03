package github

import (
	"fmt"
	"time"

	"github.com/google/go-github/github"
)

func (c *Client) GetRepoInfo(owner string, name string) (*github.Repository, error) {
	repo, resp, err := c.github.Repositories.Get(owner, name)
	if resp.StatusCode == 403 {
		resp.Body.Close()
		rateLimitUntil := c.githubRateLimitedUntil(c.github)
		fmt.Println("rate limited until", rateLimitUntil)
		time.Sleep(-1 * time.Since(rateLimitUntil))
		return nil, nil
	} else if resp.StatusCode == 404 {
		fmt.Println("404", owner, name)
		return nil, nil
	}
	return repo, err
}

func (c *Client) githubRateLimitedUntil(githubClient *github.Client) time.Time {
	rateLimits, _, err := githubClient.RateLimits()
	if err != nil {
		panic(err)
	}
	return rateLimits.Core.Reset.Time
}
