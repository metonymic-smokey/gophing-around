package main

import (
	"fmt"
	html "golang.org/x/net/html"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("no filename entered!")
		os.Exit(1)
	}

	f := os.Args[1]

	file, err := ioutil.ReadFile(f)
	if err != nil {
		log.Fatal(err)
	}

	content := string(file)
	//converting to io.Reader for tokenizer
	x := strings.NewReader(content)

	tok := html.NewTokenizer(x)
	for {
		tt := tok.Next()

		switch {
		case tt == html.ErrorToken:
			// End of the document
			return
		case tt == html.StartTagToken:
			t := tok.Token()

			//links in HTML are represented as <a href=""></a>
			isAnchor := t.Data == "a"
			if isAnchor {
				//fmt.Println("Link found!")

				for _, a := range t.Attr {
					if a.Key == "href" {
						fmt.Println("Found link:", a.Val)
						break
					}
				}
			}
		}
	}

}
