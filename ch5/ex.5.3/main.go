package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fmtlinks1: %v\n", err)
		os.Exit(1)
	}
	for _, text := range visit(nil, doc) {
		fmt.Println(text)
	}
}

func visit(texts []string, n *html.Node) []string {
	if n == nil {
		return texts
	}
	if n.Type == html.ElementNode && (n.Data == "script" || n.Data == "style") {
		return texts
	}
	if n.Type == html.TextNode {
		texts = append(texts, n.Data)
	}
	texts = visit(texts, n.FirstChild)
	texts = visit(texts, n.NextSibling)

	return texts
}
