package links

import (
	"fmt"
	"io/ioutil"
	"regexp"
)

func ParseFile(f string) [][]byte {
	content, err := ioutil.ReadFile(f)
	if err != nil {
		fmt.Println(err)
	}
	re := regexp.MustCompile(`((http|https)+\:\/\/)([\w\d-]+\.)*[\w-]+[\.\:]\w+([\/\?\=\&\#]?[\w-]+)*\/?`)
	u := re.FindAll(content, -1)
	return u
}
