package model

type AmzKeywordModel struct {
	NextPageURL string       `json:"nextPage"`
	PageIndex   string       `json:"pageIndex"`
	Keyword     string       `json:"keyword"`
	Products    []AmzProduct `json:"products"`
}

type AmzProduct struct {
	ASIN      string `json:"asin"`
	Price     string `json:"price"`
	Star      string `json:"star"`
	Title     string `json:"title"`
	Image     string `json:"img"`
	Rating    string `json:"rating"`
	Sales     string `json:"sales"`
	Prime     string `json:"prime"`
	Sponsered string `json:"sponsered"`
}
