package github

import (
	"context"

	"github.com/google/go-github/v48/github"
)

func PostComment(
	owner string,
	repo string,
	issueNumber int,
	comment string,
) (*github.Response, error) {
	ctx := context.Background()
	client := GetClient()

	_, res, err := client.Issues.CreateComment(ctx, owner, repo, issueNumber, &github.IssueComment{
		Body: &comment,
	})

	if err != nil {
		return nil, err
	}

	return res, nil
}
