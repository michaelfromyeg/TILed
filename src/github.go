package main

import (
	"context"
	"fmt"

	"github.com/google/go-github/github"

	// "github.com/libgit2/git2go/v30"
	"io/ioutil"
	"os"

	"golang.org/x/oauth2"
)

var (
	user  string
	token string
)

func init() {
	// Register environment variables
	checkVars()
}

func create() {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)
	repo := &github.Repository{
		Name:    github.String("til"),
		Private: github.Bool(false),
	}
	fmt.Println("repo", repo)
	client.Repositories.Create(ctx, "", repo)
}

func initFile() {
	err := ioutil.WriteFile("README.md", []byte("# Hello World"), 0755)
	if err != nil {
		fmt.Printf("Unable to write file: %v", err)
	}
}

func cloneRepo() {
	fmt.Println("clone")
	// repo, err := git.Clone("git@github.com:michaelfromyeg/til.git", "web", &git.CloneOptions{})
	// if err != nil {
	// 	panic(err)
	// }
}

func commitReadme() {
	fmt.Println("readme")
}

func checkVars() {
	// Get the GITHUB_USERNAME environment variable
	githubUsername, exists := os.LookupEnv("GITHUB_USERNAME")
	if exists {
		fmt.Println("Username: ", githubUsername)
	}
	user = githubUsername

	// Get the GITHUB_TOKEN environment variable
	githubToken, exists := os.LookupEnv("GITHUB_TOKEN")
	if exists {
		fmt.Println("Token: ", githubToken)
	}
	token = githubToken
}
