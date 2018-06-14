package parse

import (
	"fake-user-agent-go-ng/infra/download"
	"fmt"
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
		fmt.Println(UserAgentCom(doc))
	}
}
