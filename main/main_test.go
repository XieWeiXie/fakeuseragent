package main_test

import (
	"fmt"

	"github.com/wuxiaoxiaoshen/fakeuseragent/application"
)

func Example() {
	fakeUserAgent := application.NewFakeUserAgent(true, false)
	fmt.Println(fakeUserAgent.Random())
	fmt.Println(fakeUserAgent.Safari())
	fmt.Println(fakeUserAgent.Chrome())
	fmt.Println(fakeUserAgent.IE())
	fmt.Println(fakeUserAgent.Opera())
}
