/**
  @Author 王睿
  @Date 2020/6/16
  @Note
**/
package SearchData

import (
	"fmt"
	"github.com/gocolly/colly"
)

func Search(string2 string) {
	c := colly.NewCollector(

		colly.AllowedDomains("fund.eastmoney.com"),
	)



	c.OnHTML("#gz_gszzl", func(e *colly.HTMLElement) {

		// Print link
		fmt.Printf("Link found 实时估值: %q \n", e.Text)

	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// Start scraping on https://hackerspaces.org
	c.Visit("http://fund.eastmoney.com/"+string2+".html")
}
