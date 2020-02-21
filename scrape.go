package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/labstack/echo"
)

//CityData of homes for a given city
type cityData struct {
	//Required gorm fields
	// ID        uint       `json:"-" gorm:"primary_key"`
	// CreatedAt time.Time  `json:"-"`
	// UpdatedAt time.Time  `json:"-"`
	// DeletedAt *time.Time `json:"-" sql:"index"`
	gorm.Model
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
		//Cache responses to be nice to realestateabc :)
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
		cityValues.CityName = cityName

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
				cityValues.HomeValue = rightInt
			case "AVG PER SQ FT:":
				cityValues.AvgSqft = rightInt
			case "Property Tax:":
				cityValues.PropertyTax = rightInt
			case "Median Condo Value:":
				cityValues.MedianCondo = rightInt
			case "Median Single Family Value:":
				cityValues.SingleFamilyValue = rightInt
			case "Median 2 BD Value:":
				cityValues.TwoBDValue = rightInt
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
	//GORM Stuff
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(cityData{})

	//echo Stuff
	e := echo.New()
	//Get request with a specified id
	e.GET("/citydata/:id", func(c echo.Context) error {
		//Convert id to int
		pk, _ := strconv.Atoi(c.Param("id"))

		//Get Record from db
		record := db.First(&cityData{}, pk)

		//return record to user as json
		return c.JSON(200, record)

	})

	//episodeLinks()

	homeValues := stateView()
	fmt.Println("STATE DATA COMPILED")
	//fmt.Println(homeValues)

	for _, value := range homeValues {
		fmt.Println("-------------------")
		db.Create(&value)
		fmt.Printf("Added %s to DB\n", value.CityName)
	}

	//homeValuesJSON, _ := json.Marshal(homeValues)
	fmt.Println("-------------------")
	fmt.Println("DATA CONVERTED")
	//fmt.Println(homeValuesJSON)

	//Start Echo server and log
	e.Logger.Fatal(e.Start(":1323"))

}
