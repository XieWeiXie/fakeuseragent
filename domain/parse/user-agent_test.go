package parse

import (
	"fmt"
	"github.com/wuxiaoxiaoshen/fakeuseragent/infra/download"
	"testing"
)

func TestUserAgentCom(t *testing.T) {
	tests := []struct {
		browserType string
	}{
		{
			browserType: "Chrome",
		},
		{
			browserType: "Safari",
		},
	}

	for _, tt := range tests {
		url := fmt.Sprintf("http://useragentstring.com/pages/useragentstring.php?name=%s", tt.browserType)
		doc, _ := download.ResponseDownload(url)
		result, _ := UserAgentCom(doc)
		fmt.Println(result)
		if len(result) == 0 {
			t.Errorf("UserAgentCom() should not be nil")
		}
	}
}
