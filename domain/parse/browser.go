package parse

var (
	browsersMap = make(map[string]string)
)

func init() {
	browsersMap = map[string]string{
		"internet explorer": "internetexplorer",
		"ie":                "internetexplorer",
		"msie":              "internetexplorer",
		"edge":              "internetexplorer",
		"google":            "chrome",
		"googlechrome":      "chrome",
		"ff":                "firefox",
		"sf":                "safari",
		"op":                "opera",
	}
}

func Statistics() []string {
	browserList := make([]string, 1)
	for _, value := range browsersMap {
		browserList = append(browserList, value)
	}
	return browserList

}
