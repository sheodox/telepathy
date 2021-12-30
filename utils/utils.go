package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func PromptString(prompt string) (string, error) {
	fmt.Printf("%v ", prompt)

	reader := bufio.NewReader(os.Stdin)
	response, err := reader.ReadString('\n')

	if err != nil {
		return "", err
	}
	response = strings.TrimSuffix(response, "\n")

	return response, nil
}

func PromptStringWithFallback(prompt, fallback string) (string, error) {
	response, err := PromptString(prompt)

	if err != nil {
		return "", err
	}

	if response == "" {
		return fallback, nil
	}

	return response, nil
}
