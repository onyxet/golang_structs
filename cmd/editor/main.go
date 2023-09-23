package main

import "structs/pkg/editor"

func main() {
	editor.WriteToFileFromCMD("test")
	editor.SearchInFile("test")
}
