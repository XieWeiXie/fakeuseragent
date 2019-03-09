package application

import (
	"errors"
	"fmt"

	"github.com/wuxiaoxiaoshen/fakeuseragent/domain/global"
	"github.com/wuxiaoxiaoshen/fakeuseragent/infra/download"

	"github.com/wuxiaoxiaoshen/fakeuseragent/domain/parse"

	"math/rand"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/tidwall/gjson"
)

// FakeUserAgent: CloudFrontNetOk field should always be false
type FakeUserAgent struct {
	UserAgentStringOk bool
	CloudFrontNetOk   bool
	Cache             bool
}

var (
	ErrUserAgent = errors.New("user agent err")
)

func NewFakeUserAgent(UserAgentStringOk bool, CloudFrontNetOk bool, CacheOK bool) *FakeUserAgent {
	return &FakeUserAgent{
		UserAgentStringOk: UserAgentStringOk,
		CloudFrontNetOk:   CloudFrontNetOk,
		Cache:             CacheOK,
	}
}

// IE UserAgent
func (F *FakeUserAgent) IE() string {
	return F.common("Internet+Explorer")

}

// InternetExplorer UserAgent
func (F *FakeUserAgent) InternetExplorer() string {
	return F.IE()
}

// Msie UserAgent
func (F *FakeUserAgent) Msie() string {
	return F.IE()
}

// Chrome UserAgent
func (F *FakeUserAgent) Chrome() string {
	return F.common("Chrome")
}

// Google UserAgent
func (F *FakeUserAgent) Google() string {
	return F.Chrome()
}

// Opera UserAgent
func (F *FakeUserAgent) Opera() string {
	return F.common("Opera")
}

// Safari UserAgent
func (F *FakeUserAgent) Safari() string {
	return F.common("Safari")
}

// FireFox UserAgent
func (F *FakeUserAgent) FireFox() string {
	return F.common("Firefox")
}

// FF UserAgent
func (F *FakeUserAgent) FF() string {
	return F.FireFox()
}

// Random UserAgent
func (F *FakeUserAgent) Random() string {
	randomChoice := []string{
		"Chrome",
		"Firefox",
		"Safari",
		"Opera",
		"Internet+Explorer",
	}
	r := rand.NewSource(time.Now().UnixNano())
	random := rand.New(r)

	browserType := randomChoice[random.Intn(len(randomChoice))]
	return F.common(browserType)
}

func (F *FakeUserAgent) common(browserType string) string {
	r := rand.NewSource(time.Now().Unix())
	randomChoice := rand.New(r)
	if F.Cache {
		index := randomChoice.Intn(len(global.LOCALUSERAGENT[browserType]))
		return global.LOCALUSERAGENT[browserType][index]

	}

	var url string
	if F.UserAgentStringOk {
		url = fmt.Sprintf(global.BROWSER_BASE_PAGE, browserType)
	} else {
		url = global.CACHE_SERVER
	}

	var (
		doc *goquery.Document
		err error
	)

	doc, err = download.ResponseDownload(url)

	if err != nil {
		fmt.Println(ErrUserAgent)
		panic(ErrUserAgent)
	}
	var (
		userAgentList []string
	)

	if F.UserAgentStringOk {
		userAgentList, err = parse.UserAgentCom(doc)
		if err != nil {
			fmt.Println(ErrUserAgent)
			panic(ErrUserAgent)
		}
		return userAgentList[randomChoice.Intn(len(userAgentList))]
	}

	if F.CloudFrontNetOk {
		var userAgentResult []gjson.Result
		userAgentResult, err = parse.CloudFront(doc, browserType)
		if err != nil {
			fmt.Println(ErrUserAgent)
			panic(ErrUserAgent)
		}
		return userAgentResult[randomChoice.Intn(len(userAgentResult))].String()
	}
	return ""

}
