package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/gocolly/colly"
)

func main() {
	fName := "data.csv"
	file, err := os.Create(fName) //we create a file data.csv

	if err != nil {
		log.Fatalf("Could not create file. err: %q", err)
		return
	}

	defer file.Close() //'defer' defers the function "Close" until the oter functions are done

	writer := csv.NewWriter(file)
	defer writer.Flush()

	c := colly.NewCollector(
		colly.AllowedDomains("internshala.com"),
	)

	//Point to the web structure and specify what i  need from the page
	c.OnHTML(".internship_meta", func(e *colly.HTMLElement) {
		writer.Write([]string{
			e.ChildText("a"), //get the text from the tag a
			e.ChildText("span"),
		})
	})

	for i := 0; i < 327; i++ {
		fmt.Printf("Scraping Page: %d\n", i)
		c.Visit("https://internshala.com/internships/page-" + strconv.Itoa(i))
	}

	log.Printf("Scraping Complete\n")
	log.Panicln(c)
}
