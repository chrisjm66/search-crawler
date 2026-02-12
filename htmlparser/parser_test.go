package htmlparser

import (
	"fmt"
	"slices"
	"testing"
)

func TestLinkExtraction(t *testing.T) {
	testcases := [...]HtmlTestcase{
		{
			filePath: "./html-test-files/test1.html",
			links: []string{
				"https://something.com",
				"https://somethingelse.org",
			},
		},
		{
			filePath: "./html-test-files/test2.html",
			links: []string{
				"https://example.com",
				"https://dontforgetthehead.com",
			},
		},
	}

	for _, v := range testcases {
		reader, err := ReadHtmlFile(v.filePath)
		parsedLinks, err := ParseHtml(reader)

		fmt.Printf("Links returned from HTML parsing: %s\n", parsedLinks)
		fmt.Printf("Links expected as defined in test: %s\n", v.links)

		if err != nil {
			t.Errorf("Test failed: error %s\n", err.Error())
		}

		if len(parsedLinks) != len(v.links) {
			fmt.Printf("lengths not equal: %d %d \n", len(parsedLinks), len(v.links))
			t.Errorf("Test failed: parsed links and links table of differing length.\n")
		}

		for _, link := range parsedLinks {
			if !slices.Contains(v.links, link) {
				t.Errorf("Test failed: link added that should not be included.\n")
			}
		}
	}
}

type HtmlTestcase struct {
	filePath string
	links    []string
}
