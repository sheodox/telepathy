package server

import (
	"fmt"
	"net/http"
	"os"
	"path"

	"github.com/gin-gonic/gin"
)

var sharedDirectory string

func bootServer(dir string) error {
	sharedDirectory = dir
	router := gin.Default()
	router.GET("/files", getFiles)
	router.GET("/files/:fileName", getFileByName)

	router.Run("localhost:5676")

	return nil
}

func getFiles(c *gin.Context) {
	var files []string
	file, err := os.Open(sharedDirectory)

	if err != nil {
		getFilesError(err, c)
		return
	}

	fileList, err := file.ReadDir(-1)

	if err != nil {
		getFilesError(err, c)
		return
	}

	for _, file := range fileList {
		files = append(files, file.Name())
	}

	c.IndentedJSON(http.StatusOK, files)
}

func getFilesError(err error, c *gin.Context) {
	fmt.Printf("Error listing files %s", err)
	c.JSON(http.StatusInternalServerError, gin.H{"error": "error listing files"})
}

func getFileByName(c *gin.Context) {
	fmt.Println(path.Join(sharedDirectory, c.Param("fileName")))
	c.File(path.Join(sharedDirectory, c.Param("fileName")))
}
