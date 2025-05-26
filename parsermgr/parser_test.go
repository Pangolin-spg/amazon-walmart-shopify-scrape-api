package parsermgr

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/Pangolin-spg/amazon-walmart-shopify-scrape-api/model"
	"github.com/Pangolin-spg/amazon-walmart-shopify-scrape-api/parsermgr/amazon"
	"github.com/antchfx/htmlquery"
)

func TestAmzKeywordParser(t *testing.T) {
	p, err := NewParserManager()
	if err != nil {
		t.Fatalf("Error creating parser manager: %s\n", err.Error())
	}

	doc, err := htmlquery.LoadDoc("./input.html")
	if err != nil {
		t.Fatalf("Error loading document: %s\n", err.Error())
	}

	region, err := amazon.ParseRegion(doc)
	if err != nil {
		t.Fatalf("Error parsing region: %s\n", err.Error())
	}

	parser := p.GetAmzKeywordParser(region)
	if parser == nil {
		t.Fatal("No parser found for region: " + region)
	}

	pageIndex, err := parser.ParseCurrentPageIndex(doc)
	if err != nil {
		t.Errorf("Error parsing current page index: %s\n", err.Error())
	}

	nextPage, err := parser.ParseNextPageURL(doc)
	if err != nil {
		t.Errorf("Error parsing next page: %s\n", err.Error())
	}

	keyword, err := parser.ParseKeyword(doc)
	if err != nil {
		t.Errorf("Error parsing keyword: %s\n", err.Error())
	}

	nodes, err := parser.ParseAllProducts(doc)
	if err != nil {
		t.Fatalf("Error parsing products: %s\n", err.Error())
	}

	result := &model.AmzKeywordModel{
		NextPageURL: nextPage,
		PageIndex:   pageIndex,
		Keyword:     keyword,
		Products:    make([]model.AmzProduct, 0),
	}

	for _, node := range nodes {
		asin, err := parser.ParseASIN(node)
		if err != nil {
			t.Errorf("Error parsing asin: %s\n", err.Error())
		}

		price, err := parser.ParsePrice(node)
		if err != nil {
			t.Errorf("Error parsing price: %s\n", err.Error())
		}

		star, err := parser.ParseStar(node)
		if err != nil {
			t.Errorf("Error parsing star: %s\n", err.Error())
		}

		rating, err := parser.ParseRating(node)
		if err != nil {
			t.Errorf("Error parsing rating: %s\n", err.Error())
		}

		sponsered, _ := parser.ParseSponsered(node)

		prime, _ := parser.ParsePrime(node)

		sales, _ := parser.ParseSales(node)

		title, err := parser.ParseTitle(node)
		if err != nil {
			t.Errorf("Error parsing title: %s\n", err.Error())
		}

		img, err := parser.ParseImg(node)
		if err != nil {
			t.Errorf("Error parsing img: %s\n", err.Error())
		}

		product := model.AmzProduct{
			ASIN:      asin,
			Price:     price,
			Star:      star,
			Title:     title,
			Image:     img,
			Rating:    rating,
			Sales:     sales,
			Prime:     prime,
			Sponsered: sponsered,
		}
		result.Products = append(result.Products, product)
	}

	Printf(result)
}

func Printf(v interface{}) {
	data, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		fmt.Println("JSON序列化错误:", err)
		return
	}
	fmt.Println(string(data))
}
