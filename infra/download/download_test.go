package download

import (
	"fmt"
	"testing"
)

func TestResponseDownload(t *testing.T) {

	tests := []struct {
		url string
	}{
		{
			url: "http://d2g6u4gh6d9rq0.cloudfront.net/browsers/fake_useragent_0.1.9.json",
		},
		{
			url: "https://github.com/wuxiaoxiaoshen",
		},
	}

	for _, tt := range tests {
		_, err := ResponseDownload(tt.url)
		fmt.Println(err)
		if err != nil {
			t.Error("response download failed")
		}

	}

}
