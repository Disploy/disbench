package github

import (
	"context"

	"github.com/google/go-github/v48/github"
	"golang.org/x/oauth2"
)

var client *github.Client

func InitClient(token string) {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	client = github.NewClient(tc)

}

func GetClient() *github.Client {
	if client == nil {
		panic("github client is not initialized")
	}

	return client
}
