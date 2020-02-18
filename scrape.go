package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func realestate() {
	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting")
	})

	c.OnHTML(".px11.darkbrown.bold", func(e *colly.HTMLElement) {
		link := e.Attr("href")

		fmt.Printf("Link found: %s -> %s \n", e.Text, link)
	})

	c.Visit("https://www.realestateabc.com/home-values/search/GA")
}

func episodeLinks() {
	//Manages Network Collection Executes Callbacks
	c := colly.NewCollector()

	//On Request
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	//On Found HTML
	c.OnHTML("b a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		title := e.Attr("title")
		fmt.Printf("Link found: %q -> %s\n", title, link)
	})

	//Website to visit
	c.Visit("https://vampirediaries.fandom.com/wiki/Season_One_(Legacies)")
}

// main() contains code adapted from example found in Colly's docs:
// http://go-colly.org/docs/examples/basic/
func main() {

	//episodeLinks()

	realestate()

}
