package main

import (
	"context"
	"fmt"
	"os"

	"github.com/google/go-github/github"
	"github.com/naren-m/github-cli/services"
)

func main() {
	ctx := context.Background()
	client := github.NewClient(nil)

	githubAPI := services.NewGithub(ctx, client.Repositories)
	repo, err := githubAPI.GetPackageRepoInfo("naren-m", "dotfiles")

	if err != nil {
		fmt.Printf("Problem in getting repository information %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("%+v\n", repo)

}
