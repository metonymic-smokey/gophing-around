package main

import (
	"container/list"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func GET(urlFlag string) []string {
	resp, err := http.Get(urlFlag)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	reqURL := resp.Request.URL
	baseURL := &url.URL{
		Scheme: reqURL.Scheme,
		Host:   reqURL.Host,
	}

	base := baseURL.String()

	return filter(base, buildLinks(resp.Body, base))
}

func filter(base string, links []string) []string {
	var fl []string
	//filtered links

	for _, x := range links {
		if strings.HasPrefix(x, base) {
			fl = append(fl, x)
		}
	}

	return fl
}

func buildLinks(body io.Reader, base string) []string {

	data, _ := ioutil.ReadAll(body)

	var res []Link
	res = parse(string(data))

	var hrefs []string

	for _, link := range res {
		switch {
		case link.Href[0] == '/':
			hrefs = append(hrefs, base+link.Href)
		case strings.HasPrefix(link.Href, "http"):
			//http is a common prefix for both http and https
			hrefs = append(hrefs, link.Href)

		}
	}

	return hrefs
}

func bfs(urlFlag string, height int) []string {

	visited := make(map[string]bool)

	q := list.New()
	nq := list.New()
	nq.PushBack(urlFlag)

	for i := 0; i <= height; i++ {
		q, nq = nq, list.New()

		for temp := q.Front(); temp != nil; temp = temp.Next() {
			if visited[string(temp.Value.(string))] == true {
				continue
			}

			visited[temp.Value.(string)] = true
			for _, x := range GET(temp.Value.(string)) {
				nq.PushBack(x)
			}

		}

	}

	var vis []string
	for x, _ := range visited {
		vis = append(vis, x)
	}

	return vis
}

func main() {

	urlFlag := flag.String("url", "http://gophercises.com", "your choice of URL")
	height := flag.Int("height", 2, "depth to which tree is traversed")
	flag.Parse()
	//do not name var url - leads to clash with net/url

	pages := bfs(*urlFlag, *height)

	for _, x := range pages {
		fmt.Println(x)
	}

}
