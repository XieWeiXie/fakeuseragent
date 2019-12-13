package download

import (
	"fmt"
	"testing"

	"github.com/PuerkitoBio/goquery"
)

func TestResponseDownload(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name    string
		args    args
		want    *goquery.Document
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test for download from baidu.com",
			args: args{
				url: "https://www.baidu.com",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			re, err := ResponseDownload(tt.args.url)
			fmt.Println(re.Url)
			if err != nil {
				t.Errorf("ResponseDownload() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

		})
	}
}
