package parse

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

var browserList = []string{
	"Internet Explorer",
	"Opera",
	"Chrome",
	"Safari",
	"FireFox",
}

func UserAgentCom(doc *goquery.Document) ([]string, error) {
	doc.Find("div#liste ul li").Each(func(i int, selection *goquery.Selection) {
		userAgent := selection.Find("a").Text()
		fmt.Println(userAgent)
	})
	return []string{}, nil

}
