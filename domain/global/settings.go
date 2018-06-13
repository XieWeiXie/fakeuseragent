package global

import "fmt"

const Version = "0.1.10"

var (
	BROWSERS_STATS_PAGE = "https://www.w3schools.com/browsers/default.asp"
	BROWSER_BASE_PAGE   = "http://useragentstring.com/pages/useragentstring.php?name=%s"
	CACHE_SERVER        = "http://d2g6u4gh6d9rq0.cloudfront.net/browsers/fake_useragent_%s.json"
)

var OVERRIDES = make(map[string]string)

func init() {
	OVERRIDES = map[string]string{
		"Edge/IE": "Internet Explorer",
		"IE/Edge": "Internet Explorer",
	}
	CACHE_SERVER = fmt.Sprintf(CACHE_SERVER, Version)
}
