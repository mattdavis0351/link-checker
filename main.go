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

func main() {
	fn := files.ReadWorkspaceDir()
	fmt.Printf("the following files will be parsed:\n%s\n", fn)
	lo := links.AsListOfObjects(fn)
	fmt.Printf("the following objects have been created from the files to parse:\n%v\n", lo)
	// o := "[{some:super long}, {kinda:complex, string: of things}]"

	j, err := json.Marshal(lo)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("json data after marshalling:\n%s\n", string(j))
	err = actions.SetOutput(string(j))
	if err != nil {
		fmt.Println(err)
	}

	// fmt.Println(lo)
}
