package server

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"path"

	"github.com/gin-gonic/gin"
)

var sharedDirectory string

func bootServer(dir string) error {
	sharedDirectory = dir
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/files")
	})
	router.GET("/files", getFiles)
	router.GET("/files/:fileName", getFileByName)

	router.Run("localhost:5676")

	return nil
}

func getFiles(c *gin.Context) {
	var files []string
	f, err := os.Open(sharedDirectory)
	defer f.Close()

	if err != nil {
		getFilesError(err, c)
		return
	}

	fileList, err := f.ReadDir(-1)

	if err != nil {
		getFilesError(err, c)
		return
	}

	for _, file := range fileList {
		if !file.IsDir() {
			files = append(files, file.Name())
		}
	}

	c.IndentedJSON(http.StatusOK, files)
}

func getFilesError(err error, c *gin.Context) {
	fmt.Printf("Error listing files %s", err)
	c.JSON(http.StatusInternalServerError, gin.H{"error": "error listing files"})
}

func getFileByName(c *gin.Context) {
	fileName, err := url.QueryUnescape(c.Param("fileName"))
	if err != nil {
		return
	}

	c.File(path.Join(sharedDirectory, fileName))
}
