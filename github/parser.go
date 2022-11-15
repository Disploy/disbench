package github

import (
	"fmt"
	"strconv"
	"strings"
)

type ParsedGitHubTarget struct {
	Owner       string
	Repo        string
	IssueNumber int
}

func ParseGitHubTarget(target string) (ParsedGitHubTarget, error) {
	var parsedTarget ParsedGitHubTarget

	var parts = strings.Split(target, "/")

	if len(parts) != 2 {
		return parsedTarget, fmt.Errorf("invalid target: %s", target)
	}

	parsedTarget.Owner = parts[0]

	var repoAndIssue = strings.Split(parts[1], "#")
	parsedTarget.Repo = repoAndIssue[0]

	var issueNumber, err = strconv.Atoi(repoAndIssue[1])

	if err != nil {
		return parsedTarget, fmt.Errorf("invalid target: %s", target)
	}

	parsedTarget.IssueNumber = issueNumber

	return parsedTarget, err
}
