package digits_search

import (
	"fmt"
	"os"
	"regexp"
)

const regex = "(\\+?\\d{1,3}[-.\\s]?)?(\\()?(\\d{1,3})(?(2)\\))[-.\\s]?(\\d{3})[-.\\s]?(\\d{4})\n"

type DigitFileReader struct {
	content string
}

func (d *DigitFileReader) ReadFile(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	d.content = string(data)
	return nil
}

func (d *DigitFileReader) Search() ([]string, error) {
	r, err := regexp.Compile(regex)
	if err != nil {
		fmt.Errorf("unexpected error during regex registration %g", err)
		return nil, err
	}
	match := r.FindAllString(d.content, -1)
	return match, nil
}
