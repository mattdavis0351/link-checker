package main

import (
	"encoding/json"
	"fmt"

	"github.com/mattdavis0351/link-checker/actions"
	"github.com/mattdavis0351/link-checker/files"
	"github.com/mattdavis0351/link-checker/links"
)

// /([\w+]+\:\/\/)?([\w\d-]+\.)*[\w-]+[\.\:]\w+([\/\?\=\&\#]?[\w-]+)*\/?/gm

// TODO
// read a file line by line
// check the line against the url regex
// 	if there is a link make a GET request to the link
// 		status 200 -> link is valid... count it as a link in the repo
// 		status !200 > link is not valid, count it, report as broken
// Set the resulting output as environment variable using actions workflow commands
type urls struct {
	URLs []links.Link `json:"urls"`
}

func main() {
	fn := files.ReadWorkspaceDir()
	lo := links.AsListOfObjects(fn)

	u := urls{lo}

	j, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}

	err = actions.SetOutput(string(j))
	if err != nil {
		fmt.Println(err)
	}
}
