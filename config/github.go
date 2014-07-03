package config

import (
	"os"

	"code.google.com/p/goauth2/oauth"
	upGithub "github.com/dickeyxxx/updeps/github"
	"github.com/google/go-github/github"
)

func Github() *upGithub.Client {
	token := os.Getenv("GITHUB_ACCESS_TOKEN")
	if token != "" {
		t := &oauth.Transport{
			Token: &oauth.Token{AccessToken: os.Getenv("GITHUB_ACCESS_TOKEN")},
		}
		return upGithub.NewClient(github.NewClient(t.Client()))
	} else {
		return upGithub.NewClient(github.NewClient(nil))
	}
}
