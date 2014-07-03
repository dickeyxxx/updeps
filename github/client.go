package github

import "github.com/google/go-github/github"

type Client struct {
	github *github.Client
}

func NewClient(client *github.Client) *Client {
	return &Client{client}
}
