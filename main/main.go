package main

import (
	"fmt"

	"github.com/wuxiaoxiaoshen/fakeuseragent/application"
)

func main() {
	fakeUserAgent := application.NewFakeUserAgent(false)
	fmt.Println(fakeUserAgent.Random())
	fmt.Println(fakeUserAgent.Safari())
	fmt.Println(fakeUserAgent.Chrome())
	fmt.Println(fakeUserAgent.IE())
	fmt.Println(fakeUserAgent.Opera())
}
