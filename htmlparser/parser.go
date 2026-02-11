package htmlparser

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	"golang.org/x/net/html"
)

func ParseHtml(reader io.Reader) ([]string, error) {
	node, err := html.Parse(reader)
	links := []string{}

	if err != nil {
		return nil, err
	}

	links = traverse(node, links)

	return links, nil
}

func traverse(node *html.Node, links []string) []string {
	if node.Type == html.ElementNode {
		switch node.Data {
		case "a":
			link, err := extractLink(node)
			fmt.Printf("Link found: %s\n", link)
			if err != nil {
				fmt.Printf("Error: %s\n", err)
				break
			}

			links = append(links, link)
		}
	}

	for i := node.FirstChild; i != nil; i = i.NextSibling {
		if i.Type == html.ElementNode  && i.Data != "head"{
			fmt.Println(i.Data)
			links = traverse(i, links)
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
