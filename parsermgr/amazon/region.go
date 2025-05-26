package amazon

import (
	"github.com/Pangolin-spg/amazon-walmart-shopify-scrape-api/errors"
	"github.com/Pangolin-spg/amazon-walmart-shopify-scrape-api/utils"
	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
)

const (
	US = "en-us"
	UK = "en-gb"
	DE = "de-de"
	FR = "fr-fr"
	CA = "en-ca"
)

const (
	US_PREFIX = "https://www.amazon.com"
	UK_PREFIX = "https://www.amazon.co.uk"
	DE_PREFIX = "https://www.amazon.de"
	FR_PREFIX = "https://www.amazon.fr"
	CA_PREFIX = "https://www.amazon.ca"
)

var regionMap = map[string]string{
	US: US_PREFIX,
	UK: UK_PREFIX,
	DE: DE_PREFIX,
	FR: FR_PREFIX,
	CA: CA_PREFIX,
}

func GetPrefix(region string) string {
	if prefix, ok := regionMap[region]; ok {
		return prefix
	}
	return US_PREFIX
}

func ParseRegion(doc *html.Node) (string, error) {
	langExpr := "/html/@lang"
	langNodes, err := utils.FindNodes(doc, langExpr, false)
	if err != nil {
		return "unknown", err
	}

	lang := htmlquery.SelectAttr(langNodes[0], "lang")
	if utils.StringIsEmpty(lang) {
		return "unknown", errors.ErrorNotFoundLanguage
	}
	return lang, nil
}
