package main

import (
	"fmt"
	"fund/SearchData"

	"github.com/gocolly/colly"
)

func main() {
	// Instantiate default collector
	c := colly.NewCollector(
		// Visit only domains: hackerspaces.org, wiki.hackerspaces.org
		colly.AllowedDomains("hackerspaces.org", "wiki.hackerspaces.org","fund.eastmoney.com"),
	)

	// On every a element which has href attribute call callback
	//c.OnHTML("a[href]", func(e *colly.HTMLElement) {
	//	link := e.Attr("href")
	//	// Print link
	//	fmt.Printf("Link found: %q -> %s\n", e.Text, link)
	//	// Visit link found on page
	//	// Only those links are visited which are in AllowedDomains
	//	//c.Visit(e.Request.AbsoluteURL(link))
	//})

	c.OnHTML("#gz_gszzl", func(e *colly.HTMLElement) {

		// Print link
		fmt.Printf("Link found 实时估值: %q \n", e.Text)

	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// Start scraping on https://hackerspaces.org
	//c.Visit("https://hackerspaces.org/")
	c.Visit("http://fund.eastmoney.com/004851.html")
	SearchData.Search("000529")

}
