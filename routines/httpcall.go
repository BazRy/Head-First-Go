package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type Page struct {
	URL  string
	Size int
}

func main() {
	pages := make(chan Page)
	startInMillis := time.Now().UnixMilli()
	urls := []string{"https://www.mozilla.org/en-GB/firefox/", "https://golang.org", "https://google.com", "https://bbc.com", "https://sky.com"}

	for _, url := range urls {
		go responseSize(url, pages)
	}
	for i := 0; i < len(urls); i++ {
		page := <-pages
		fmt.Printf("%s: %d\n", page.URL, page.Size)
	}

	endInMillis := time.Now().UnixMilli()
	println("Time taken: ", endInMillis-startInMillis, "ms")
}

func responseSize(url string, myChannel chan Page) {
	fmt.Println("Getting", url)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	myChannel <- Page{URL: url, Size: len(body)}
}
