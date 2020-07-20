package git

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

var (
	user  string
	token string
)

func createGithubRepository() {
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

func initLocalFile(directory string) {
	s := fmt.Sprintf("%v/README.md", directory)
	err := ioutil.WriteFile(s, []byte("# This file was created in the command line"), 0755)
	CheckIfError(err)
}

func deleteLocalFile() {
	err := os.Remove("README.md")
	CheckIfError(err)
}

func updateFile(directory string, message string, url string) {
	s := fmt.Sprintf("%v/README.md", directory)

	file, err := os.OpenFile(s, os.O_APPEND|os.O_WRONLY, 0644)
	CheckIfError(err)
	defer file.Close()

	newRow := createMarkdownRow(message, url)
	_, err = file.WriteString(newRow)
	_, err = file.WriteString("\n")
	CheckIfError(err)
}

func createMarkdownRow(message string, url string) string {
	t := time.Now()
	t2 := t.Format("2006-01-02")
	row := fmt.Sprintf("| %v | %v | %v |", t2, message, url)
	return row
}

func addFile(directory string) {
	// Opens an already existing repository.
	fmt.Println("Opening repo")
	r, err := git.PlainOpen(directory)
	CheckIfError(err)

	fmt.Println("Opening tree")
	w, err := r.Worktree()
	CheckIfError(err)

	fmt.Println("add file")

	s := "README.md"
	_, err = w.Add(s)
	CheckIfError(err)
}

func cloneGithubRepository(directory string) {
	repo := fmt.Sprintf("github.com/%v/til.git", user)
	url := fmt.Sprintf("https://%s:%s@%s", user, token, repo)
	// url := fmt.Sprintf("https://github.com/%s/til", user)
	_, err := git.PlainClone(directory, false, &git.CloneOptions{
		URL:               url,
		RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
		Progress:          os.Stdout,
	})
	CheckIfError(err)
}

func commitFile(name string, email string, directory string) {
	// Opens an already existing repository.
	fmt.Println("Opening repo")
	r, err := git.PlainOpen(directory)
	CheckIfError(err)

	fmt.Println("Opening tree")
	w, err := r.Worktree()
	CheckIfError(err)

	// We can verify the current status of the worktree using the method Status.
	fmt.Println("Get status")
	status, err := w.Status()
	CheckIfError(err)
	fmt.Println(status)

	fmt.Println("Trying to commit")
	commit, err := w.Commit("Commit created by TIL CLI", &git.CommitOptions{
		Author: &object.Signature{
			Name:  name,
			Email: email,
			When:  time.Now(),
		},
	})
	CheckIfError(err)

	obj, err := r.CommitObject(commit)
	CheckIfError(err)

	fmt.Println(obj)
}

func pushChanges(directory string) {
	r, err := git.PlainOpen(directory)
	CheckIfError(err)

	err = r.Push(&git.PushOptions{})
	CheckIfError(err)
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
