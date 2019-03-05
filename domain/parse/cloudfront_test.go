package parse

import (
	"fmt"
	"testing"

	"github.com/wuxiaoxiaoshen/fakeuseragent/infra/download"
)

func TestCloudFront(t *testing.T) {
	doc, _ := download.ResponseDownload("http://d2g6u4gh6d9rq0.cloudfront.net/browsers/fake_useragent_0.1.9.json")
	browsersType := []string{
		"chrome", "opera", "firefox", "safari", "internetexplorer",
	}
	for _, browser := range browsersType {
		rt, err := CloudFront(doc, browser)
		if err != nil {

		}
		fmt.Println(rt)
	}

}
