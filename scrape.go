package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

// main() contains code adapted from example found in Colly's docs:
// http://go-colly.org/docs/examples/basic/
func main() {
	// Instantiate default collector
	c := colly.NewCollector()

	// On every a element which has href attribute call callback
	c.OnHTML("b a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		title := e.Attr("title")
		fmt.Printf("Link found: %q -> %s\n", title, link)

	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.Visit("https://vampirediaries.fandom.com/wiki/Season_One_(Legacies)")
}
