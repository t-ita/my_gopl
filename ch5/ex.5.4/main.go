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
	var linksmap = make(map[string][]string)
	visit(linksmap, doc)
	for name, links := range linksmap {
		fmt.Printf("%s ------------------------------\n", name)
		for _, link := range links {
			fmt.Printf("\t%s\n", link)
		}
	}
}

func visit(linksmap map[string][]string, n *html.Node) {
	if n == nil {
		return
	}
	if n.Type == html.ElementNode {
		for _, a := range n.Attr {
			if a.Key == "href" {
				linksmap[n.Data] = append(linksmap[n.Data], a.Val)
			}
		}
	}
	visit(linksmap, n.FirstChild)
	visit(linksmap, n.NextSibling)
}
