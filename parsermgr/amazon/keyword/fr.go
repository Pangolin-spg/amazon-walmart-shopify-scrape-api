package keyword

import (
	"strings"

	"github.com/Pangolin-spg/amazon-walmart-shopify-scrape-api/errors"
	"github.com/Pangolin-spg/amazon-walmart-shopify-scrape-api/utils"
	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
)

type FRKeywordParser struct{}

func NewFRKeywordParser() *FRKeywordParser {
	return &FRKeywordParser{}
}

func (p *FRKeywordParser) ParseAllProducts(doc *html.Node) ([]*html.Node, error) {
	expr := "//div[@class and @data-asin and string-length(@data-asin) > 0 and @data-index and @data-uuid]"

	nodes, err := utils.FindNodes(doc, expr, true)
	if err != nil {
		return nil, err
	}
	return nodes, nil
}

func (p *FRKeywordParser) parseCurrentPageIndexV1(doc *html.Node) (string, error) {
	expr := `//span[contains(@class, 's-pagination-selected')]/text()`

	nodes, err := utils.FindNodes(doc, expr, true)
	if err != nil || len(nodes) == 0 {
		return "", err
	}
	return strings.TrimSpace(nodes[0].Data), nil
}

func (p *FRKeywordParser) parseCurrentPageIndexV2(doc *html.Node) (string, error) {
	expr := `//span[contains(@aria-label, 'Page actuelle')]/text()`

	nodes, err := utils.FindNodes(doc, expr, true)
	if err != nil || len(nodes) == 0 {
		return "", err
	}
	return strings.TrimSpace(nodes[0].Data), nil
}

func (p *FRKeywordParser) ParseCurrentPageIndex(doc *html.Node) (string, error) {
	pageIndex, err := p.parseCurrentPageIndexV1(doc)
	if err == nil && pageIndex != "" {
		return pageIndex, err
	}

	pageIndex, err = p.parseCurrentPageIndexV2(doc)
	if err == nil && pageIndex != "" {
		return pageIndex, err
	}
	return "unknown", errors.ErrorNotFoundPageIndex
}

func (p *FRKeywordParser) parseNextPageURLV1(doc *html.Node) (string, error) {
	expr := `//a[contains(@class, 's-pagination-next')]`

	nodes, err := utils.FindNodes(doc, expr, false)
	if err != nil || len(nodes) == 0 {
		return "", err
	}
	return htmlquery.SelectAttr(nodes[0], "href"), nil
}

func (p *FRKeywordParser) parseNextPageURLV2(doc *html.Node) (string, error) {
	expr := `//a[contains(@aria-label, 'Accéder à la page suivant')]`

	nodes, err := utils.FindNodes(doc, expr, false)
	if err != nil || len(nodes) == 0 {
		return "", err
	}
	return htmlquery.SelectAttr(nodes[0], "href"), nil
}

func (p *FRKeywordParser) ParseNextPageURL(doc *html.Node) (string, error) {
	nextPage, err := p.parseNextPageURLV1(doc)
	if err == nil && nextPage != "" {
		return nextPage, err
	}

	nextPage, err = p.parseNextPageURLV2(doc)
	if err == nil && nextPage != "" {
		return nextPage, err
	}
	return "unknown", errors.ErrorNotFoundNextPage
}

func (p *FRKeywordParser) ParseKeyword(doc *html.Node) (string, error) {
	expr := `//input[@id='twotabsearchtextbox']/@value`

	nodes, err := utils.FindNodes(doc, expr, true)
	if err != nil {
		return "unknown", err
	}

	keyword := htmlquery.SelectAttr(nodes[0], "value")
	if utils.StringIsEmpty(keyword) {
		return "unknown", errors.ErrorNotFoundImgURL
	}
	return keyword, nil
}

func (p *FRKeywordParser) ParseASIN(node *html.Node) (string, error) {
	expr := `@data-asin`

	nodes, err := utils.FindNodes(node, expr, true)
	if err != nil {
		return "unknown", err
	}
	return htmlquery.SelectAttr(nodes[0], "data-asin"), nil
}

func (p *FRKeywordParser) ParsePrice(node *html.Node) (string, error) {
	expr := `//div//span[@class='a-price']/span/text()`

	nodes, err := utils.FindNodes(node, expr, true)
	if err != nil {
		return "unknown", err
	}

	price := strings.TrimSpace(nodes[0].Data)
	if utils.StringIsEmpty(price) {
		return "unknown", nil
	}
	return price, nil
}

func (p *FRKeywordParser) ParseStar(node *html.Node) (string, error) {
	expr := `//div//span[contains(@aria-label,'sur 5')]`

	nodes, err := utils.FindNodes(node, expr, true)
	if err != nil {
		return "unknown", err
	}
	stars := htmlquery.SelectAttr(nodes[0], "aria-label")
	star := utils.FormatNumberEuro(strings.Split(stars, " ")[0])
	if utils.StringIsEmpty(star) {
		return "unknown", nil
	}
	return star, nil
}

func (p *FRKeywordParser) ParseRating(node *html.Node) (string, error) {
	expr := `//span[contains(@aria-label, 'évaluations')]`

	nodes, err := utils.FindNodes(node, expr, true)
	if err != nil {
		return "unknown", err
	}
	return utils.FormatRating(htmlquery.SelectAttr(nodes[0], "aria-label")), nil
}

func (p *FRKeywordParser) ParseSponsered(node *html.Node) (string, error) {
	expr := `//div//span[text()='Sponsorisé']`

	_, err := utils.FindNodes(node, expr, true)
	if err != nil {
		return "0", err
	}
	return "1", nil
}

func (p *FRKeywordParser) ParsePrime(node *html.Node) (string, error) {
	expr := `//div//i[@aria-label="Amazon Prime"]`

	_, err := utils.FindNodes(node, expr, true)
	if err != nil {
		return "false", err
	}
	return "true", nil
}

func (p *FRKeywordParser) ParseSales(node *html.Node) (string, error) {
	expr := `//div//span[contains(text(), "achetés au cours du mois dernier")]/text()`

	nodes, err := utils.FindNodes(node, expr, true)
	if err != nil {
		return "unknown", err
	}
	sales := strings.Trim(nodes[0].Data, "achetés au cours du mois dernier")
	return utils.FormatNumber(sales), nil
}

func (p *FRKeywordParser) ParseImg(node *html.Node) (string, error) {
	expr := `//div//img[contains(@class,"image")]/@src`

	nodes, err := utils.FindNodes(node, expr, true)
	if err != nil {
		return "unknown", err
	}
	return htmlquery.SelectAttr(nodes[0], "src"), nil
}

func (p *FRKeywordParser) ParseTitle(node *html.Node) (string, error) {
	expr := `//div//span[contains(@class, "text-normal")]/text()`

	nodes, err := utils.FindNodes(node, expr, true)
	if err != nil {
		return "unknown", err
	}
	return utils.FormatTitle(nodes[0].Data), nil
}
