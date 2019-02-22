package tapula_test

import (
	"fmt"
	"testing"

	"github.com/antchfx/htmlquery"
	"github.com/johnmcdnl/tapula"
)

func Test_FindWebElement(t *testing.T) {
	xpath := "//div[contains(@class, 'something')]//a[2]"
	link := tapula.FindWebElement(testdata(), xpath)
	fmt.Printf("%s(%s)\n", htmlquery.InnerText(link), htmlquery.SelectAttr(link, "href"))
}

func Test_FindWebElements(t *testing.T) {

	xpath := `//a[contains(@class, 'something')]/ancestor::tr//div[contains(@class, 'somethingelse')]//a`
	links := tapula.FindWebElements(testdata(), xpath)

	for i, link := range links {
		fmt.Printf("%d %s(%s)\n", i, htmlquery.InnerText(link), htmlquery.SelectAttr(link, "href"))
	}
}

func Test_GetLinkTarget(t *testing.T) {
	xpath := `//a[contains(@class, 'abc')]/ancestor::tr//div[contains(@class, 'def')]//a`
	links := tapula.FindWebElements(testdata(), xpath)

	for _, link := range links {
		fmt.Println(tapula.GetLinkTarget(link))
	}

}
