package parse

import (
	"fmt"
	"testing"

	"github.com/wuxiaoxiaoshen/fakeuseragent/infra/download"
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
		download.ResponseDownload(url)
	}
}
