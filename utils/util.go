package utils

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
)

func FindNodes(doc *html.Node, expr string, multi bool) ([]*html.Node, error) {
	nodes, err := htmlquery.QueryAll(doc, expr)
	if err != nil {
		return nil, fmt.Errorf("'%v' error, %v", expr, err)
	}

	if len(nodes) == 0 {
		return nil, fmt.Errorf("'%v' error, no nodes selected", expr)
	}

	if len(nodes) != 1 && !multi {
		return nil, fmt.Errorf("'%v' error, %v nodes selected", expr, len(nodes))
	}
	return nodes, err
}

func FormatNumber(s string) string {
	s = strings.ReplaceAll(s, "k", "000")
	s = strings.ReplaceAll(s, "K", "000")
	s = strings.ReplaceAll(s, ",", "")
	return strings.ReplaceAll(s, "+", "")
}

func FormatNumberEuro(old string) string {
	old = strings.ReplaceAll(old, ".", "")
	old = strings.ReplaceAll(old, ",", ".")
	old = strings.ReplaceAll(old, "k", "000")
	return strings.ReplaceAll(old, "K", "000")
}

func FormatTitle(s string) string {
	return strings.Replace(strings.Replace(strings.TrimSpace(s), "'", " ", -1), "\"", " ", -1)
}

func FormatRating(s string) string {
	lastIndex := strings.LastIndexFunc(s, unicode.IsDigit)
	if lastIndex == -1 {
		return s
	}
	str := strings.ReplaceAll(s[0:lastIndex+1], ",", "")
	str = strings.ReplaceAll(str, ".", "")
	return strings.ReplaceAll(str, " ", "")
}
