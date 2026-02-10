package htmlparser

import (
	"fmt"
	"io"
	"os"

	"golang.org/x/net/html"
)

func ParseHtml(reader io.Reader) {
	token := html.NewTokenizer(reader)

	for {
		tt := token.Next()

		if tt == html.ErrorToken {
			break;
		}

		fmt.Println(token.TagName())
	}
}

func ReadHtmlFile(filepath string) {
	file, err := os.ReadFile(filepath)	
	// TODO	
}
