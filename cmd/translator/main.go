package main

import (
	"net/http"
	"structs/pkg/translator"
)

func main() {
	router := translator.Router()
	http.ListenAndServe(":8080", router)
}
