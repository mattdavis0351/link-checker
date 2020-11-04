package files

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func ReadWorkspaceDir() []string {
	files, err := ioutil.ReadDir(os.Getenv("GITHUB_WORKSPACE"))
	var fileNames []string
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		fileNames = append(fileNames, file.Name())

		fmt.Println(file.Name())

	}
	return fileNames

}
