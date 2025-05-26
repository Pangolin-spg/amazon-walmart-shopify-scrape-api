package parsermgr

import (
	"github.com/Pangolin-spg/amazon-walmart-shopify-scrape-api/parsermgr/amazon"
	"github.com/Pangolin-spg/amazon-walmart-shopify-scrape-api/parsermgr/amazon/keyword"
)

const (
	AmzKeyword = "amzKeyword"
)

// ParserManager is responsible for parsing the messages from the message queue.
type ParserManager struct {
	amzKeywordParserMap map[string]AmzKeywordParser
}

// NewParserManager creates a new ParserManager.
func NewParserManager() (*ParserManager, error) {
	pm := &ParserManager{
		amzKeywordParserMap: make(map[string]AmzKeywordParser),
	}

	// Register all parsers.
	pm.registerParsers()
	return pm, nil
}

func (pm *ParserManager) registerParser(region string, parser interface{}) {
	switch p := parser.(type) {

	case AmzKeywordParser:
		pm.amzKeywordParserMap[region] = p

	default:
		return
	}
}

func (pm *ParserManager) GetAmzKeywordParser(region string) AmzKeywordParser {
	return pm.amzKeywordParserMap[region]
}

func (pm *ParserManager) registerParsers() {
	// Register keyword parsers.
	pm.registerParser(amazon.US, keyword.NewUSKeywordParser())
	pm.registerParser(amazon.UK, keyword.NewUKKeywordParser())
	pm.registerParser(amazon.DE, keyword.NewDEKeywordParser())
	pm.registerParser(amazon.FR, keyword.NewFRKeywordParser())
	pm.registerParser(amazon.CA, keyword.NewCAKeywordParser())
}
