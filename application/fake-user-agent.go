package application

import (
	"errors"
	"fmt"
	"log"

	"github.com/wuxiaoxiaoshen/fakeuseragent/domain/global"
	"github.com/wuxiaoxiaoshen/fakeuseragent/infra/download"

	"github.com/wuxiaoxiaoshen/fakeuseragent/domain/parse"

	"math/rand"
	"time"

	"github.com/PuerkitoBio/goquery"
)

// FakeUserAgent ...
type FakeUserAgent struct {
	UserAgentStringOk bool // should be remove
	Cache             bool
}

var (
	ErrUserAgent = errors.New("user agent err")
)

func NewFakeUserAgent(UserAgentStringOk bool, CacheOK bool) *FakeUserAgent {
	return &FakeUserAgent{
		UserAgentStringOk: UserAgentStringOk,
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
	var browserType string
	browserType = randomChoice[random.Intn(len(randomChoice))]

	return F.common(browserType)
}

func (F *FakeUserAgent) common(browserType string) string {
	if !F.valid() {
		log.Println("UserAgentStringOk or CacheOk should be true")
		return "None"
	}
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
	return ""

}

func (F *FakeUserAgent) valid() bool {
	// UserAgentStringOk
	// Cache
	// 两个参数必须其一为 true
	if !(F.UserAgentStringOk || F.Cache) {
		return false
	}
	return true
}
