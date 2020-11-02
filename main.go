package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"net/http"
)


// /([\w+]+\:\/\/)?([\w\d-]+\.)*[\w-]+[\.\:]\w+([\/\?\=\&\#]?[\w-]+)*\/?/gm

// TODO
// read a file line by line
// check the line against the url regex
// 	if there is a link make a GET request to the link
// 		status 200 -> link is valid... count it as a link in the repo
// 		status !200 > link is not valid, count it, report as broken
// Write results to an issue on GitHub
var good, bad []bytes
func main(){
	
	f, err := ioutil.ReadFile("test-txt-file.txt")
	if err != nil{
		fmt.Println("something went terribly wrong")
	}
	// string from []byte of test-txt-file.txt
	// s := string(f)

	re := regexp.MustCompile(`([\w+]+\:\/\/)?([\w\d-]+\.)*[\w-]+[\.\:]\w+([\/\?\=\&\#]?[\w-]+)*\/?`)
	u := re.FindAll([]byte(f), -1)

	for _, v := range u{
		resp, err := http.Get(string(v))
		defer resp.Body.Close()
		// http status code do not count as err
		if err != nil{
			fmt.Println(err)
		}

		if resp.StatusCode != 200{
			fmt.Printf("Not 200, got: %d", resp.StatusCode)
		}
		// fmt.Printf("%s\n",v)
		fmt.Println(resp.StatusCode)
	}
	// fmt.Println(s)

}