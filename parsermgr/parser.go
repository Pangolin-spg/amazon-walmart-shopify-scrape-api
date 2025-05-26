package parsermgr

import "golang.org/x/net/html"

type AmzKeywordParser interface {
	// ParseAllProducts parses all products from the given HTML document.
	ParseAllProducts(doc *html.Node) ([]*html.Node, error)

	// ParseCurrentPageIndex parses the current page index from the given HTML document.
	ParseCurrentPageIndex(doc *html.Node) (string, error)

	// ParseNextPageURL parses the next page url from the given HTML document.
	ParseNextPageURL(doc *html.Node) (string, error)

	// ParseKeyword parses the keyword from the given HTML document.
	ParseKeyword(doc *html.Node) (string, error)

	// ParseASIN parses the ASIN from the given HTML node.
	ParseASIN(node *html.Node) (string, error)

	// ParsePrice parses the price from the given HTML node.
	ParsePrice(node *html.Node) (string, error)

	// ParseStar parses the star from the given HTML node.
	ParseStar(node *html.Node) (string, error)

	// ParseRating parses the rating from the given HTML node.
	ParseRating(node *html.Node) (string, error)

	// ParseSponsered parses the sponsered from the given HTML node.
	ParseSponsered(node *html.Node) (string, error)

	// ParsePrime parses the prime from the given HTML node.
	ParsePrime(node *html.Node) (string, error)

	// ParseSales parses the sales from the given HTML node.
	ParseSales(node *html.Node) (string, error)

	// ParseImg parses the img from the given HTML node.
	ParseImg(node *html.Node) (string, error)

	// ParseTitle parses the title from the given HTML node.
	ParseTitle(node *html.Node) (string, error)
}
