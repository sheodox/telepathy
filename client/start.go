package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
	"strconv"
	"time"

	"github.com/sheodox/telepathy/utils"
)

func Start() error {
	hostname, err := utils.PromptString("Enter the IP of a computer running Telepathy server:")

	if err != nil {
		return err
	}

	filesUrl := fmt.Sprintf("http://%v:5676/files", hostname)

	for {
		files, err := fetchList(filesUrl)

		if err != nil {
			return err
		}

		if len(files) == 0 {
			return errors.New("That computer doesn't have any files to share!")
		}

		fmt.Println()
		for i, fileName := range files {
			fmt.Printf("%v. %v\n", i+1, fileName)
		}

		choiceStr, err := utils.PromptString("Enter the number of the file you wish to download (r to refresh, q to quit):")

		if err != nil {
			return err
		}

		switch choiceStr {
		case "q":
			return nil
		case "r":
			// do nothing, let the loop restart
		default:
			choice, err := strconv.Atoi(choiceStr)
			// prompt uses numbers starting at 1, change back to an index
			choice--

			if err != nil || choice < 0 || choice >= len(files) {
				fmt.Printf("Number must be between 1 and %v.\n", len(files))
				pause()
				continue
			}

			download(filesUrl, files[choice])
		}
	}
}

func download(filesUrl, file string) error {
	startTime := time.Now()
	url := fmt.Sprintf("%v/%v", filesUrl, url.QueryEscape(file))
	homeDir, err := os.UserHomeDir()

	if err != nil {
		return err
	}

	downloadDest := path.Join(homeDir, "/Downloads", file)
	fmt.Printf("Downloading %v to %v\n", file, downloadDest)
	resp, err := http.Get(url)

	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(downloadDest)

	if err != nil {
		return err
	}

	defer out.Close()

	_, err = io.Copy(out, resp.Body)

	fmt.Printf("Done in %v\n", time.Now().Sub(startTime))
	pause()

	return err
}

func pause() {
	utils.PromptString("Press enter to continue...")
}

func fetchList(url string) ([]string, error) {
	resp, err := http.Get(url)

	var files []string

	if err != nil {
		return files, err
	}

	json.NewDecoder(resp.Body).Decode(&files)

	return files, nil
}
