package link

import (
	"fmt"
	"golang.org/x/net/html"
	"io"
)

// this struct defines a link
type Link struct {
	href string
	text string
}

// parse a reader to generate links
func Parse(r io.Reader) ([]Link, error) {
	root, err := html.Parse(r)
	if err != nil {
		return nil, err
	}
	traverse(root, "")
	return nil, nil
}

func traverse(node *html.Node, padding string) {
	fmt.Println(padding, node.Data)
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		traverse(c, padding+" ")
	}

}
