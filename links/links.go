package links

import (
	"fmt"
	"net/http"
)

type Link struct {
	URL        string `json:"url"`
	FileName   string `json:"file_name"`
	StatusCode int    `json:"status_code"`
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
		foundURLs := ParseFile(file)
		linkList := urlsAsList(file, foundURLs)
		for _, linkItem := range linkList {
			l = append(l, linkItem)
		}
	}
	return l
}
