package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"

	"github.com/gocolly/colly"
)

type Item struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
}

func main() {

	allItems := make([]Item, 0)

	c := colly.NewCollector(
		colly.AllowedDomains("factretriever.com", "www.factretriever.com"),
	)

	//Point to the web structure and specify what i  need from the page
	c.OnHTML(".factsList li", func(e *colly.HTMLElement) {
		itemId, err := strconv.Atoi(e.Attr("id"))
		if err != nil {
			log.Println("Could not get id")
		}

		itemText := e.Text

		item := Item{
			ID:          itemId,
			Description: itemText,
		}

		allItems = append(allItems, item)

	})

	c.OnRequest(func(request *colly.Request) {
		fmt.Println("Visiting: ", request.URL.String())
	})

	c.Visit("https://www.factretriever.com/rhino-facts")

	log.Printf("Scraping Complete\n")
	log.Println(c)
	writeJSON(allItems)
}

func writeJSON(data []Item) {
	file, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		log.Println("Unable to create JSON")
		return
	}

	_ = ioutil.WriteFile("dataInJSON.json", file, 0644)
}
