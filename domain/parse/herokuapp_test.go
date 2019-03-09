package parse

import (
	"strings"
	"testing"

	"github.com/wuxiaoxiaoshen/fakeuseragent/infra/download"

	"github.com/PuerkitoBio/goquery"
)

func TestHerokuApp(t *testing.T) {
	type args struct {
		doc         *goquery.Document
		browserType string
	}
	doc, _ := download.ResponseDownload("https://fake-useragent.herokuapp.com/browsers/0.1.8")
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "get user agent from herokuapp.com",
			args: args{
				doc:         doc,
				browserType: "chrome",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//fmt.Println(tt.args.doc.Html())
			got, err := HerokuApp(tt.args.doc, tt.args.browserType)
			if err != nil {
				t.Errorf("HerokuApp() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !strings.Contains(got, "Mozilla") {
				t.Errorf("HerokuApp() != %s", got)
			}

		})
	}
}
