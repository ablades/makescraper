package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

//CityData of homes for a given city
type cityData struct {
	CityName          string `json:"city name"`
	AvgSqft           int    `json:"avg sqft"`
	TwoBDValue        int    `json:"two bedroom value"`
	PropertyTax       int    `json:"property tax"`
	HomeValue         int    `json:"zillow home value"`
	MedianCondo       int    `json:"condo value"`
	SingleFamilyValue int    `json:"single family value"`
}

//Looks at indivdiual cities
func cityView(cityName string, link string) cityData {

	//Values for a given city
	var cityValues cityData

	c := colly.NewCollector(
		// Cache responses to prevent multiple download of pages
		colly.CacheDir("./cache"),
	)

	//Link to a specific city
	cityLink := fmt.Sprintf("https://www.realestateabc.com%s", link)

	//On page request
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("-------------------")
		fmt.Println("Visiting: " + cityName)
	})

	//Grab property table
	c.OnHTML("#propertydetails", func(e *colly.HTMLElement) {

		//Values for a given city
		cityValues.cityName = cityName

		//Loop through table data,
		e.ForEach("#propertydetails tr", func(_ int, el *colly.HTMLElement) {
			fmt.Println(el.ChildText(".subjectmenutblleft") + el.ChildText(".subjectmenutblright"))

			//Text from both columns
			leftText := el.ChildText(".subjectmenutblleft")
			rightText := el.ChildText(".subjectmenutblright")

			// Remove $ and , from text and convert to an int
			rightText = strings.ReplaceAll(rightText, ",", "")
			rightInt, _ := strconv.Atoi(rightText[1:])

			//Set value based on table data
			switch leftText {
			case "Zillow Home Value Index":
				cityValues.homeValue = rightInt
			case "AVG PER SQ FT:":
				cityValues.avgSqft = rightInt
			case "Property Tax:":
				cityValues.propertyTax = rightInt
			case "Median Condo Value:":
				cityValues.medianCondo = rightInt
			case "Median Single Family Value:":
				cityValues.singleFamilyValue = rightInt
			case "Median 2 BD Value:":
				cityValues.twoBDValue = rightInt
			}
		})
	})

	c.Visit(cityLink)

	fmt.Println(cityValues)
	fmt.Println("-------------------")

	return cityValues
}

//Lots at the current state
func stateView() []cityData {

	//List of all cities values
	var valuesList []cityData

	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting")
	})

	//
	c.OnHTML(".px11.darkbrown.bold", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		fmt.Printf("Link found: %s -> %s \n", e.Text, link)

		//Visit each city in a state
		cityData := cityView(e.Text, link)

		//Add city's data to slice
		valuesList = append(valuesList, cityData)

	})

	c.Visit("https://www.realestateabc.com/home-values/search/GA")

	return valuesList
}

func main() {

	//episodeLinks()

	homeValues := stateView()
	fmt.Println("-------------------")
	fmt.Println("STATE DATA COMPILED")
	print(homeValues)

	//List of all cities values

}
