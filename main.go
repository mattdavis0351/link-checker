package main

import (
	"fmt"

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
	// var l []links.Link
	fn := files.ReadWorkspaceDir()
	// for _, f := range fn {
	// 	fmt.Printf("checking links in %s", f)
	// 	u := links.ParseFile(f)
	// 	ll := links.ListLinks(f, u)
	// 	for _, li := range ll {
	// 		l = append(l, li)
	// 	}
	// }
	lo := links.AsObjects(fn)

	fmt.Println(lo)
}
