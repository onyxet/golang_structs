package main

import (
	"flag"
	"log"
	"os"
	"regexp"
)

const numRegex = "\\(?\\d{3}\\)?[-.\\s]?\\d{3}[-.\\s]?\\d{4}"
const wordRegex = "Лес[яії]?\\s+Україн[а-яіїє]*"

// numSearch searches for any form of phone number in given text
func numSearch(filename string) {
	r, err := regexp.Compile(numRegex)
	if err != nil {
		log.Fatal(err)
	}
	fileData, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	numbers := r.FindAllString(string(fileData), -1)
	for _, number := range numbers {
		println("Match:", number)
	}
}

// wordSearch searches for any form of "Леся Українка" in given text
func wordSearch(filename string) {
	r, err := regexp.Compile(wordRegex)
	if err != nil {
		log.Fatal(err)
	}
	fileData, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	numbers := r.FindAllString(string(fileData), -1)
	for _, text := range numbers {
		println("Match text:", text)
	}
}

func main() {
	nums := flag.Bool("nums", false, "Display phone numbers in a file")
	words := flag.Bool("words", false, "search for \"Леся Українка\" in a file")
	filename := flag.String("file", "default.txt", "file to search")

	flag.Parse()

	if *nums {
		numSearch(*filename)
	} else if *words {
		wordSearch(*filename)
	} else {
		log.Println("No valid search option selected.")
	}
}
