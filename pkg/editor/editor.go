package editor

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"structs/pkg/errors"
)

func WriteToFileFromCMD(filename string) {
	reader := bufio.NewReader(os.Stdin)
	textByte := make([]byte, 0)
	fmt.Println("Please enter the text you want to write to the file. Type \"exit\" when you are done.")
	for {
		fmt.Print("> ")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSuffix(text, "\n")
		if text == "exit" {
			break
		}
		text = text + "\n"
		textByte = append(textByte, []byte(text)...)
	}
	err := os.WriteFile(filename+".txt", textByte, 0644)
	errors.Check(err)
}

func SearchInFile(filename string) {
	// Checking file existence
	_, err := os.Stat(filename + ".txt")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	// Opening the file
	f, err := os.Open(filename + ".txt")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer f.Close()

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Enter search query. Type \"exit\" when you are done.")
		fmt.Print("> ")
		searchQuery, err := reader.ReadString('\n')
		searchQuery = strings.TrimSuffix(searchQuery, "\n")
		if searchQuery == "exit" {
			break
		}
		errors.Check(err)
		found := false
		scanner := bufio.NewScanner(f)
		lineInFile := 1
		// Scanning through each line of the file
		for scanner.Scan() {
			if strings.Contains(scanner.Text(), searchQuery) {
				fmt.Printf("Line: %v\nText: %v\n", lineInFile, scanner.Text())
				found = true
			}
			lineInFile++
		}

		f.Seek(0, 0)

		if !found {
			fmt.Println("Text was not found")
		}
	}
}
