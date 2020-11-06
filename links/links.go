package links

import (
	"fmt"
	"net/http"
)

type Link struct {
	URL        string
	FileName   string
	StatusCode int
}

func urlsAsList(n string, u [][]byte) []Link {
	var a []Link
	for _, v := range u {
		l := Link{
			URL:      string(v),
			FileName: n,
		}

		resp, err := http.Get(l.URL)
		if err != nil {
			fmt.Println(err)
		}

		defer resp.Body.Close()
		l.StatusCode = resp.StatusCode
		a = append(a, l)

	}
	return a
}

func AsListOfObjects(fileNames []string) []Link {
	var l []Link
	for _, file := range fileNames {
		fmt.Printf("checking links in %s", file)
		foundURLs := ParseFile(file)
		linkList := urlsAsList(file, foundURLs)
		for _, linkItem := range linkList {
			l = append(l, linkItem)
		}
	}
	return l
}
