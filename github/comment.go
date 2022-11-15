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
) {
	ctx := context.Background()
	client := GetClient()

	client.Issues.CreateComment(ctx, owner, repo, issueNumber, &github.IssueComment{
		Body: &comment,
	})
}
