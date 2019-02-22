package tapula

import (
	"strings"

	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
)

type WebElementLink struct {
	LinkText    string
	LinkAddress string
}

func FindWebElement(body, xpath string) *html.Node {
	return htmlquery.FindOne(parseDoc(body), xpath)
}

func FindWebElements(body, xpath string) []*html.Node {
	return htmlquery.Find(parseDoc(body), xpath)
}

func GetLinkTarget(link *html.Node) string {
	return htmlquery.SelectAttr(link, "href")
}

func parseDoc(body string) *html.Node {
	doc, err := htmlquery.Parse(strings.NewReader(body))
	if err != nil {
		panic(err)
	}
	return doc
}
