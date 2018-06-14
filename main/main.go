package main

import (
	"fake-user-agent-go-ng/application"
	"fmt"
)

func main() {
	fakeUserAgent := application.NewFakeUserAgent(true, false, false)
	fmt.Println(fakeUserAgent.Random())
	fmt.Println(fakeUserAgent.Safari())
	fmt.Println(fakeUserAgent.Chrome())
	fmt.Println(fakeUserAgent.IE())
	fmt.Println(fakeUserAgent.Opera())
}
