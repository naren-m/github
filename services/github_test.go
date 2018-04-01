package services

import (
	"context"
	"fmt"
	"testing"

	"github.com/google/go-github/github"
)

func Test_get_repositories(t *testing.T) {
	client := github.NewClient(nil)
	ctx := context.Background()
	userName := "naren-m"
	orgs, _, err := client.Organizations.List(ctx, userName, nil)

	fmt.Println(orgs, err)

	// List all repositories for a user
	opt := &github.RepositoryListOptions{}
	repos, _, err := client.Repositories.List(ctx, userName, opt)

	for r := 0; r < len(repos); r++ {
		fmt.Println(*repos[r].CloneURL)
	}
	// fmt.Println(repos)
}
