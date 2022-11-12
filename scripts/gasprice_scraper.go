package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly"
)

func main() {
	fName := "../raw_data/wa_gasprices.csv"
	file, err := os.Create(fName)
	if err != nil {
		log.Fatalf("Cannot create file %q: %s\n", fName, err)
		return
	}
	defer file.Close()
	writer := csv.NewWriter(file)

	// defer writer.Flush()

	// city_names := []string{}
	c := colly.NewCollector()
	// c := colly.NewCollector(colly.Debugger(&debug.LogDebugger{}))

	// Writer headers
	writer.Write([]string{"City", "Regular", "Mid", "Premium", "Disel"})

	city_names := make([]string, 0)
	rows := [][]string{}
	counter := 0

	// Write gas prices per city names
	c.OnHTML("div.accordion-prices", func(e *colly.HTMLElement) {

		e.ForEach("h3", func(_ int, el *colly.HTMLElement) {
			city_names = append(city_names, el.Text)
		})

		e.ForEach("table.table-mob", func(_ int, el2 *colly.HTMLElement) {
			e.ForEach("tr", func(_ int, el3 *colly.HTMLElement) {
				if (el3.ChildText("td:nth-child(1)") == "Current Avg.") &&
					(counter < len(city_names)) {
					temp := make([]string, 0)
					temp = append(temp, city_names[counter])
					temp = append(temp, el3.ChildText("td:nth-child(2)"))
					temp = append(temp, el3.ChildText("td:nth-child(3)"))
					temp = append(temp, el3.ChildText("td:nth-child(4)"))
					temp = append(temp, el3.ChildText("td:nth-child(5)"))

					rows = append(rows, temp)
					counter++
				}
			})
		})

		fmt.Println(rows)
		err = writer.WriteAll(rows) // calls Flush internally

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Scrapping Complete")
	})

	// c.OnHTML("div.accordion-prices table.table-mob", func(e *colly.HTMLElement) {
	// 	e.ForEach("tr", func(_ int, el2 *colly.HTMLElement) {
	// 		if el2.ChildText("td:nth-child(1)") == "Current Avg." {
	// 			temp := make([]string, 0)
	// 			temp = append(temp, city_names[counter])
	// 			temp = append(temp, el2.ChildText("td:nth-child(2)"))
	// 			temp = append(temp, el2.ChildText("td:nth-child(3)"))
	// 			temp = append(temp, el2.ChildText("td:nth-child(4)"))
	// 			temp = append(temp, el2.ChildText("td:nth-child(5)"))

	// 			rows = append(rows, temp)
	// 			counter++
	// 		}
	// 	})

	// 	fmt.Println(rows)
	// 	err = writer.WriteAll(rows) // calls Flush internally

	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	fmt.Println("Scrapping Complete")
	// })

	c.OnResponse(func(r *colly.Response) {
		fmt.Println(r.StatusCode)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit("https://gasprices.aaa.com/?state=WA")
}
