package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
)

func main() {
	var queryFunc func(query, resource string)
	if args.urlMode {
		queryFunc = queryURL
	} else {
		queryFunc = queryFile
	}

	if len(args.resources) == 0 {
		if args.urlMode {
			sc := bufio.NewScanner(os.Stdin)
			for sc.Scan() {
				res := sc.Text()
				if len(res) != 0 {
					queryFunc(args.query, res)
				}
			}
		} else {
			doQuery(args.query, os.Stdin)
		}
	} else {
		for _, resource := range args.resources {
			queryFunc(args.query, resource)
		}
	}
}

func queryURL(query string, url string) {
	res, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Printf("status code error: %d %s\n", res.StatusCode, res.Status)
		return
	}

	doQuery(query, res.Body)
}

func queryFile(query, filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Println(err)
		return
	}

	doQuery(query, file)
}

func doQuery(query string, r io.Reader) {
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		log.Println(err)
		return
	}

	doc.Find(query).Each(func(i int, s *goquery.Selection) {
		for _, node := range s.Nodes {
			// node.Type == html.ElementNode
			err := html.Render(os.Stdout, node)
			if err != nil {
				log.Println(err)
			} else {
				fmt.Println()
			}
		}
	})
}
