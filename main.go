package golang_structs

import "regexp"

const numRegex = "\\(?(\\+?\\d{1,3})?\\)?[\\s.-]?(\\d{3})[\\s.-]?(\\d{3})[\\s.-]?(\\d{4})\n"

func numSearch(text string) int {
	r, _ := regexp.Compile(numRegex)
	return 1
}
