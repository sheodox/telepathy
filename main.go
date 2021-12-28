package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/sheodox/telepathy/server"
)

func main() {
	telepathyType, err := promptTelepathyType()

	if err != nil {
		fmt.Println(err)
	}

	if telepathyType == "s" {
		server.Start()
	}

	if telepathyType == "c" {

	}
}

func promptTelepathyType() (string, error) {
	fmt.Print("Do you want to start Telepathy as a file sharing client or server? [c/s]: ")

	reader := bufio.NewReader(os.Stdin)
	desiredType, err := reader.ReadString('\n')

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
