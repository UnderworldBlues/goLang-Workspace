package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	//importing colly
	"github.com/gocolly/colly"
)

type Product struct {
	name  string
	price string
	image string
	url   string
}

func main() {
	// initializing a new collector instance
	collector := colly.NewCollector()
	// initalizing a new slice of product data
	pokeProducts := make([]Product, 0, 700)

	collector.Visit("https://scrapeme.live/shop/")

	collector.OnError(func(response *colly.Response, err error) {
		fmt.Println("Seems like something went wrong:\nERROR: ", err)
	})

	// iterating over the list of html product elements
	collector.OnHTML("li.product", func(elemento *colly.HTMLElement) {

		var currentItem Product

		// scraping the data of interest
		currentItem.url = elemento.ChildAttr("a", "href")
		currentItem.image = elemento.ChildAttr("img", "src")
		currentItem.name = elemento.ChildText("h2")
		currentItem.price = elemento.ChildText(".price")
		// adding the product instance with scraped data to the slice of products
		pokeProducts = append(pokeProducts, currentItem)
	})

	collector.OnScraped(func(response *colly.Response) {
		fmt.Println("Scrapping done: ", response.Request.URL)
	})

	// saving the scrapped data into a csv file
	file, err := os.Create("pokeProducts.csv")
	if err != nil {
		log.Fatal("Error creating file: ", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	headers := []string{"Name", "Price", "Image", "URL"}
	writer.Write(headers)

	for _, product := range pokeProducts {
		data := []string{product.name, product.price, product.image, product.url}
		writer.Write(data)
	}

}
