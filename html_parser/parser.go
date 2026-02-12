package html_parser

import (
	"bytes"
	"errors"
	"fmt"
	"golang.org/x/net/html"
	"io"
	"os"
	"strings"
)

func ParseHtml(reader io.Reader) ([]string, error) {
	node, err := html.Parse(reader)
	links := []string{}

	if err != nil {
		return nil, err
	}

	links = parseLinks(node, links)

	return links, nil
}

func parseLinks(document *html.Node, links []string) ([]string) {
	for i := range document.Descendants() {
		if i.Type == html.ElementNode {
			switch i.Data {
			case "a":
				link, err := extractLink(i)
				fmt.Printf("Link found: %s\n", link)
				if err != nil {
					fmt.Printf("Error: %s\n", err)
					break
				}

				links = append(links, link)
			// we can add other stuff in the future like images, paragraphs here
			}
		}
	}

	return links
}

func ReadHtmlFile(filepath string) (io.Reader, error) {
	file, err := os.ReadFile(filepath)

	if err != nil {
		fmt.Printf("Error: %s", err.Error())
		return nil, err
	}

	return bytes.NewReader(file), nil
}

func extractLink(node *html.Node) (string, error) {
	for _, attribute := range node.Attr {
		if attribute.Key == "href" && strings.Contains(attribute.Val, "http") {
			return attribute.Val, nil
		}
	}

	return "", errors.New("No link found.")
}
