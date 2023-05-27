package lib

import (
	"regexp"
	"io/ioutil"
	"net/http"
)

type Website struct {
	Url		string		`json:"url"`
	Links	[]string	`json:"links"`
}

func Fetch(urls []string) <-chan Website {
	ch := make(chan Website)

	for _, url := range urls {
		go fetchUrl(url, ch)
	}

	return ch 
}

func cleanUrl(url string) string {
	size := len(url)
	return url[1:size-1]
}

func fetchUrl(url string, ch chan Website) {
	resp, _ := http.Get(url)
	html, _ := ioutil.ReadAll(resp.Body)
	expr := "\"https://.*?\""
	regx, _ := regexp.Compile(expr)
	matches := regx.FindAllString(string(html), -1)

	for index, match := range matches {
		matches[index] = cleanUrl(match)
	}
	
	ch <- Website{url, matches}
}