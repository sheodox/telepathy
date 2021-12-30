package main

import (
	"fmt"
	"strings"

	"github.com/sheodox/telepathy/client"
	"github.com/sheodox/telepathy/server"
	"github.com/sheodox/telepathy/utils"
)

func main() {
	telepathyType, err := promptTelepathyType()

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println()

	if telepathyType == "s" {
		err = server.Start()
	}

	if telepathyType == "c" {
		err = client.Start()
	}

	if err != nil {
		fmt.Println(err)
	}
}

func promptTelepathyType() (string, error) {
	desiredType, err := utils.PromptString("Do you want to start Telepathy as a file sharing client or server? [c/s]")

	if err != nil {
		return "", err
	}

	desiredType = strings.TrimSuffix(strings.ToLower(desiredType), "\n")

	if desiredType != "c" && desiredType != "s" {
		fmt.Printf("Invalid choice, must be c (client) or s (server), got %q\n", desiredType)
		return promptTelepathyType()
	}

	return desiredType, nil
}
