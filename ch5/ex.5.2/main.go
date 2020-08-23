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
	var enameCount = make(map[string]int)
	visit(enameCount, doc)
	for ename, count := range enameCount {
		fmt.Printf("%s\t%d\n", ename, count)
	}
}

func visit(enameCount map[string]int, n *html.Node) {
	if n == nil {
		return
	}
	if n.Type == html.ElementNode {
		enameCount[n.Data]++
	}
	visit(enameCount, n.FirstChild)
	visit(enameCount, n.NextSibling)
}
