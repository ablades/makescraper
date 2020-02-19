package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

//HomeValues of homes for a given city
type HomeValues struct {
	cityName   string
	avgSqft    int
	twoBDValue int
	valueIndex int
}

//Looks at indivdiual cities
func cityView(link string) {
	c := colly.NewCollector(
		// Cache responses to prevent multiple download of pages
		colly.CacheDir("./cache"),
	)

	//Link to a specific city
	cityLink := fmt.Sprintf("https://www.realestateabc.com%s", link)

	//On page request
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting" + cityLink)
	})

	//Grab property table
	c.OnHTML("#propertydetails", func(e *colly.HTMLElement) {

		fmt.Println("Table " + e.Text)
		// 	//Iterate over table data
		// 	e.ForEach("#propertydetails tr", func(_ int, el *colly.HTMLElement) {
		// 		switch el.ChildText("td:subjectmenutblleft") {
		// 		case "Language":
		// 			course.Language = el.ChildText("td:nth-child(2)")
		// 		case "Level":
		// 			course.Level = el.ChildText("td:nth-child(2)")
		// 		case "Commitment":
		// 			course.Commitment = el.ChildText("td:nth-child(2)")
		// 		case "How To Pass":
		// 			course.HowToPass = el.ChildText("td:nth-child(2)")
		// 		case "User Ratings":
		// 			course.Rating = el.ChildText("td:nth-child(2) div:nth-of-type(2)")
		// 		}
		// 	})
		// 	courses = append(courses, course)

	})

	c.Visit(cityLink)
}

//Lots at the current state
func stateView() {
	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting")
	})

	c.OnHTML(".px11.darkbrown.bold", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		fmt.Printf("Link found: %s -> %s \n", e.Text, link)
		cityView(link)

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

	stateView()

}
