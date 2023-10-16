package main

import (
	"net/http"
	"structs/pkg/taskAPI"
)

func main() {

	router := taskAPI.Router()
	//log.Info().Msg("The service is ready to listen and serve.")
	http.ListenAndServe(":8080", router)
}
