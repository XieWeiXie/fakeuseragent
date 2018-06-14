package parse

import (
	"errors"

	"github.com/PuerkitoBio/goquery"
	"github.com/tidwall/gjson"
)

var (
	ErrArray   = errors.New("array err")
	ErrInvalid = errors.New("invalid json")
)

func CloudFront(doc *goquery.Document, browserType string) ([]gjson.Result, error) {

	if !gjson.Valid(doc.Text()) {
		return nil, ErrInvalid
	}

	jsonResult := gjson.Parse(doc.Text())
	browserUserAgent := jsonResult.Get("browsers." + browserType)
	browserUserAgentOk := browserUserAgent.IsArray()
	if !browserUserAgentOk {
		return nil, ErrArray
	}
	return browserUserAgent.Array(), nil

}
