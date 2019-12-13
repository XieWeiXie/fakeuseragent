package download

import (
	"log"
	"net/http"

	"errors"

	"github.com/PuerkitoBio/goquery"
)

var (
	ErrRequest  = errors.New("request err")
	ErrResponse = errors.New("response err")
)

// Get html resource by http
// return *goquery.Document and error
func ResponseDownload(url string) (*goquery.Document, error) {
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Println(err)
		return nil, ErrRequest
	}

	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2228.0 Safari/537.36")
	client := http.DefaultClient

	response, err := client.Do(request)
	if err != nil {
		log.Println(err)
		return nil, ErrResponse
	}
	return goquery.NewDocumentFromReader(response.Body)
}
