package main

import (
	"net/http"
	"structs/pkg/weather"
)

func main() {
	router := weather.Router()
	http.ListenAndServe(":8080", router)
}
