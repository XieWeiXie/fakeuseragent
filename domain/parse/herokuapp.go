package parse

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/tidwall/gjson"
)

func HerokuApp(doc *goquery.Document, browserType string) (string, error) {
	jsonData := gjson.Parse(doc.Text())
	oneBrowser := jsonData.Get("browsers").Get(browserType)

	var results string
	if !oneBrowser.Exists() {
		return results, fmt.Errorf("no exists browser %s", browserType)
	}

	array := oneBrowser.Array()
	rand.Seed(time.Now().UnixNano())

	return array[rand.Intn(len(array))].String(), nil
}
