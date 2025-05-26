package errors

import "fmt"

var (
	ErrorEmptyPage         = fmt.Errorf("page is empty")
	ErrorNotFoundASIN      = fmt.Errorf("not found asin")
	ErrorNotFoundPrice     = fmt.Errorf("not found price")
	ErrorNotFoundStar      = fmt.Errorf("not found star")
	ErrorNotFoundRating    = fmt.Errorf("not found rating")
	ErrorNotFoundRank      = fmt.Errorf("not found rank")
	ErrorNotFoundImage     = fmt.Errorf("not found image")
	ErrorNotFoundLink      = fmt.Errorf("not found link")
	ErrorNotFoundTitle     = fmt.Errorf("not found title")
	ErrorNotFoundLanguage  = fmt.Errorf("not found lang")
	ErrorNotFoundNextPage  = fmt.Errorf("not found next page")
	ErrorNotFoundPageIndex = fmt.Errorf("not found page index")
	ErrorNotFoundImgURL    = fmt.Errorf("not found img url")
)
