package main

import (
	"encoding/csv"
	"log"
	"os"

	"github.com/gocolly/colly"
)

func main() {
	fName := "ids.csv"
	fName2 := "texts.csv"
	file, err := os.Create(fName) //we create a file data.csv
	file2, err := os.Create(fName2)

	if err != nil {
		log.Fatalf("Could not create file. err: %q", err)
		return
	}

	defer file.Close() //'defer' defers the function "Close" until the oter functions are done

	defer file2.Close()

	writer := csv.NewWriter(file)
	writer2 := csv.NewWriter(file2)

	defer writer.Flush()
	defer writer2.Flush()

	c := colly.NewCollector(
		colly.AllowedDomains("factretriever.com", "www.factretriever.com"),
	)

	//Point to the web structure and specify what i  need from the page
	c.OnHTML(".factsList li", func(e *colly.HTMLElement) {
		writer.Write([]string{
			e.Attr("id"), //get the text from the tag stan
		})
	})

	c2 := colly.NewCollector(
		colly.AllowedDomains("factretriever.com", "www.factretriever.com"),
	)
	c2.OnHTML(".factsList li", func(el *colly.HTMLElement) {
		writer2.Write([]string{
			el.Text, //get the text from the tag stan
		})
	})

	c.Visit("https://www.factretriever.com/rhino-facts")
	c2.Visit("https://www.factretriever.com/rhino-facts")

	log.Printf("Scraping Complete\n")
	log.Println(c)
	log.Println(c2)
}
