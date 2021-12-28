package server

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Start() error {
	sharedDirectory, err := promptSharedDirectory()

	if err != nil {
		return err
	}

	return bootServer(sharedDirectory)
}

func promptSharedDirectory() (string, error) {
	sharedDirectory, err := readString("What directory do you want to share?", "")

	if err != nil {
		return "", err
	}

	dirExists, dirExistsErr := exists(sharedDirectory)

	if dirExists {
		return sharedDirectory, nil
	}

	if dirExistsErr != nil {
		return "", dirExistsErr
	}

	shouldCreate, err := readString("That directory doesn't exist, create it? [Y/n] ", "y")
	shouldCreate = strings.ToLower(shouldCreate)

	if shouldCreate == "y" {
		err := os.MkdirAll(sharedDirectory, 0700)

		if err != nil {
			return "", err
		}

		return sharedDirectory, nil
	}

	// non-existant directory entered, but user doesn't want to create it, ask again
	return promptSharedDirectory()
}

func readString(prompt, fallback string) (string, error) {
	fmt.Printf("%v ", prompt)

	reader := bufio.NewReader(os.Stdin)
	response, err := reader.ReadString('\n')

	if err != nil {
		return "", err
	}
	response = strings.TrimSuffix(response, "\n")

	if response == "" {
		return fallback, nil
	}

	return response, nil
}

// exists returns whether the given file or directory exists
func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
