package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		words, images, err := CountWordsAndImages(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "findlinks2: %v\n", err)
			continue
		}
		fmt.Printf("words: %d, images: %d\n", words, images)
	}
}

func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}
	words, images = countWordsAndImages(doc)
	return
}

func countWordsAndImages(n *html.Node) (words, images int) {
	if n == nil {
		return
	}
	if n.Type == html.ElementNode && (n.Data == "script" || n.Data == "style") {
		return
	}
	if n.Type == html.ElementNode && (n.Data == "img") {
		images += 1
	}
	if n.Type == html.TextNode {
		words += len(strings.Split(n.Data, " "))
	}
	var c_words, c_images int
	c_words, c_images = countWordsAndImages(n.FirstChild)
	words += c_words
	images += c_images
	c_words, c_images = countWordsAndImages(n.NextSibling)
	words += c_words
	images += c_images
	return
}
