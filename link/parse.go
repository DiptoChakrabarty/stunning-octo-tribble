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
	var links []Link
	root, err := html.Parse(r)
	if err != nil {
		return nil, err
	}
	nodes := linktraverse(root)
	for _, node := range nodes {
		links = append(links, buildLink(node))
	}
	//traverse(root, "")
	return links, nil
}

func linktraverse(node *html.Node) []*html.Node {
	if node.Type == html.ElementNode && node.Data == "a" {
		return []*html.Node{node}
	}
	var result []*html.Node
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		result = append(result, linktraverse(c)...)
	}
	return result
}

func buildLink(node *html.Node) Link {
	var result Link
	for _, attr := range node.Attr {
		if attr.Key == "href" {
			result.href = attr.Val
			break
		}
	}
	return result
}

func traverse(node *html.Node, padding string) {
	element := node.Data
	if node.Type == html.ElementNode {
		element = "<" + element + ">"
	}
	fmt.Println(padding, element)
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		traverse(c, padding+" ")
	}

}
