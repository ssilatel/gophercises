package link

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strings"

	"golang.org/x/net/html"
)

type Link struct {
	Href string
	Text string
}

func parseText(n *html.Node) string {
	if n.Type == html.TextNode {
		return n.Data
	}
	if n.Type != html.ElementNode {
		return ""
	}
	var text string
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		text += parseText(c)
	}
	return strings.Join(strings.Fields(text), " ")
}

func findNodes(n *html.Node) []*html.Node {
	if n.Type == html.ElementNode && n.Data == "a" {
		return []*html.Node{n}
	}
	var nodes []*html.Node
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		nodes = append(nodes, findNodes(c)...)
	}
	return nodes
}

func newLink(n *html.Node) Link {
	var l Link
	for _, attr := range n.Attr {
		if attr.Key == "href" {
			l.Href = attr.Val
			break
		}
	}
	l.Text = parseText(n)
	return l
}

func FindLinks(rootNode *html.Node) []Link {
	nodes := findNodes(rootNode)
	var links []Link
	for _, node := range nodes {
		links = append(links, newLink(node))
	}
	return links
}

func main() {
	filename := flag.String("f", "ex1.html", "HTML file to parse")
	flag.Parse()

	htmlFile, err := ioutil.ReadFile(*filename)
	if err != nil {
		panic(err)
	}

	doc, err := html.Parse(strings.NewReader(string(htmlFile)))
	if err != nil {
		panic(err)
	}

	nodes := findNodes(doc)
	var links []Link
	for _, node := range nodes {
		links = append(links, newLink(node))
	}

	fmt.Println(links)
}
