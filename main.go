package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/mattdavis0351/link-checker/actions"
	"github.com/mattdavis0351/link-checker/files"
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
var ignoreExtenstions string = actions.GetInput("ignored_extensions")
var includeExtenstions string = actions.GetInput("included_extensions")

func main() {
	ie := strings.Split(ignoreExtenstions, ",")
	iex := strings.Split(includeExtenstions, ",")
	fn := files.ReadWorkspaceDir()
	fmt.Println(fn)
	fmt.Printf("ignored extensions are:\n%s", ignoreExtenstions)
	fmt.Printf("length of ignore extensions slice: %d\n", len(ie))
	fmt.Printf("length of included extensions slice: %d\n", len(iex))

}
