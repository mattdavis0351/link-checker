package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// /([\w+]+\:\/\/)?([\w\d-]+\.)*[\w-]+[\.\:]\w+([\/\?\=\&\#]?[\w-]+)*\/?/gm

// TODO
// read a file line by line
// check the line against the url regex
// 	if there is a link make a GET request to the link
// 		status 200 -> link is valid... count it as a link in the repo
// 		status !200 > link is not valid, count it, report as broken
// Set the resulting output as environment variable using actions workflow commands
var ws string = os.Getenv("GITHUB_WORKSPACE")

func main() {

	files, err := ioutil.ReadDir(ws)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {

		fmt.Println(file.Name())

	}

}
