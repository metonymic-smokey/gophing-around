package main

import (
	html "golang.org/x/net/html"
	"log"
	"strings"
)

type Link struct {
	Href string
	Text string
}

func text(n *html.Node) string {
	if n.Type == html.TextNode {
		return n.Data
	}
	if n.Type != html.ElementNode {
		return ""
	}
	var ret string

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ret += text(c)
	}

	return ret

}

func parse(s string) []Link {
	var links []Link

	doc, err := html.Parse(strings.NewReader(s))
	if err != nil {
		log.Fatal(err)
	}

	var f func(*html.Node)

	f = func(n *html.Node) {
		var l Link
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					l.Href = a.Val
					b := text(n)
					b = strings.TrimSpace(b)
					l.Text = b
					links = append(links, l)
					break

				}
			}

		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}

	f(doc)

	return links
}
