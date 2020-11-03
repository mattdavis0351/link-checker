package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

// /([\w+]+\:\/\/)?([\w\d-]+\.)*[\w-]+[\.\:]\w+([\/\?\=\&\#]?[\w-]+)*\/?/gm

// TODO
// read a file line by line
// check the line against the url regex
// 	if there is a link make a GET request to the link
// 		status 200 -> link is valid... count it as a link in the repo
// 		status !200 > link is not valid, count it, report as broken
// Write results to an issue on GitHub
var good, bad []string

func parseFileForLinks(f string) [][]byte {
	content, err := ioutil.ReadFile(f)
	if err != nil {
		fmt.Println("something went terribly wrong")
	}
	re := regexp.MustCompile(`([\w+]+\:\/\/)?([\w\d-]+\.)*[\w-]+[\.\:]\w+([\/\?\=\&\#]?[\w-]+)*\/?`)
	u := re.FindAll(content, -1)
	return u
}

func main() {

	urls := parseFileForLinks("test-txt-file.txt")

	for _, v := range urls {
		sv := string(v)
		resp, err := http.Get(sv)

		// http status code do not count as err
		if err != nil {
			fmt.Println(err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != 200 {
			// fmt.Printf("Not 200, got: %d", resp.StatusCode)
			bad = append(bad, sv)
		} else {
			good = append(good, sv)
		}
		// fmt.Printf("%s\n",v)

	}
	// fmt.Println(s)
	fmt.Println(good)
	fmt.Println(bad)

}
