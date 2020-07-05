package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	flag "github.com/ogier/pflag"
)

var (
	message string
)

func main() {
	// Parse flags
	flag.Parse()

	// Register environment variables
	checkVars()

	// If user does not supply flags, print usage
	if flag.NFlag() == 0 || flag.NFlag() >= 2 {
		printUsage()
	}

	create()
	initFile()
}

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found. Please see the usage guide for more information!")
		os.Exit(1)
	}

	flag.StringVarP(&message, "message", "m", "", "Set the message to append to the TIL readme.")
}

func printUsage() {
	fmt.Printf("Usage: %s [options]\n", os.Args[0])
	fmt.Println("Options:")
	flag.PrintDefaults()
	os.Exit(1)
}
