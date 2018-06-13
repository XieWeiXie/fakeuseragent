package parse

import (
	"errors"

	"github.com/PuerkitoBio/goquery"
	"github.com/tidwall/gjson"
)

var (
	ErrArray = errors.New("array err")
)

func CloudFront(doc *goquery.Document, browserType string) ([]gjson.Result, error) {
	jsonResult := gjson.Parse(doc.Text())
	browserUserAgent := jsonResult.Get("browsers." + browserType)
	browserUserAgentOk := browserUserAgent.IsArray()
	if !browserUserAgentOk {
		return []gjson.Result{}, ErrArray
	}
	return browserUserAgent.Array(), nil

}
