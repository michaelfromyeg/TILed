package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	flag "github.com/ogier/pflag"
)

var (
	message   string
	url       string
	directory string
	name      string
	email     string
)

func main() {
	// Parse flags
	flag.Parse()

	// Register environment variables
	checkVars()

	// If user does not supply flags, print usage
	if flag.NFlag() == 0 || flag.NFlag() >= 6 {
		printUsage()
	}

	// createGithubRepository() -- TODO: implement error handling
	// cloneGithubRepository(directory)
	// initLocalFile(directory)
	updateFile(directory, message, url)
	addFile(directory)
	commitFile(name, email, directory)
	pushChanges(directory)
	// deleteLocalFile()
}

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found. Please see the usage guide for more information!")
		os.Exit(1)
	}

	flag.StringVarP(&message, "message", "m", "", "Set the message to append to the TIL readme.")
	flag.StringVarP(&url, "url", "u", "", "Set the URL to append to the TIL readme.")
	flag.StringVarP(&directory, "directory", "d", "", "Set a given directory to clone the repository to. Use '.' if you'd like to clone to your current folder.")
	flag.StringVarP(&name, "name", "n", "", "Set name")
	flag.StringVarP(&email, "email", "e", "", "Set email")
}

func printUsage() {
	fmt.Printf("Usage: %s [options]\n", os.Args[0])
	fmt.Println("Options:")
	flag.PrintDefaults()
	os.Exit(1)
}
