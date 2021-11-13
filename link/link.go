package link

import (
	"io"
	"strings"

	"golang.org/x/net/html"
)

type Link struct {
	Href string
	Text string
}

// GetLinks will take io.Reader and returns all the links in HTML page in the Link data structure format.
func GetLinks(reader io.Reader) ([]Link, error) {
	var links []Link

	// html.Parse will return the *html.node type doc
	doc, err := html.Parse(reader)
	if err != nil {
		return links, err
	}

	var parser func(*html.Node)
	parser = func(n *html.Node) {

		// For the Link the node.Data will be equal to "a"
		// the node.Attr will contains the attributes of the elementNode(In our case "a" node).
		// We are going to take only the element nodes and ignore everything else
		// node.Attr is array of Attribute (Attr      []Attribute)
		if n.Type == html.ElementNode && n.Data == "a" {
			links = append(links, buildLink(n))
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			parser(c)
		}
	}
	parser(doc)
	return links, nil
}

// buildLink will extract the 'href' and 'text' and creates the 'Link'
func buildLink(node *html.Node) Link {
	var link Link

	for _, a := range node.Attr {
		if a.Key == "href" {
			link.Href = a.Val
		}
	}
	link.Text = getText(node, "")
	return link
}

// Parse the Text from the html.Node
// This will trim all extra whitespaces
func getText(node *html.Node, text string) string {

	// Link Text is part of the "TextNode"
	if node.Type == html.TextNode {
		if len(text) > 0 && !strings.HasSuffix(text, " ") {
			text = text + " "
		}
		text = text + strings.TrimSpace(node.Data)
	}
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		text = getText(c, text)
	}
	return text
}
