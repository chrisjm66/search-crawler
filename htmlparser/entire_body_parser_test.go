package htmlparser

import (
	"fmt"
	"testing"
)

func TestLinkExtraction(t *testing.T) {
	testcases := [...]HtmlTestcase{
		{
			filePath: "html-parser/html-test-files/html1.html",
			links: []string{
				"http://something.com",
				"https://somethingelse.com",
			},
		},
		{
			filePath: "html-parser/html-test-files/html2.html",
			links: []string{
				"https://example.com",
			},
		},
	}

	fmt.Print(testcases)
	t.Skip("Not implemented")
}

type HtmlTestcase struct {
	filePath  string
	links []string
}
