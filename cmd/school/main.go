package main

import (
	"github.com/rs/zerolog/log"
	"net/http"
	"structs/pkg/schoolClass"
)

func main() {

	router := schoolClass.Router()
	log.Info().Msg("The service is ready to listen and serve.")
	http.ListenAndServe(":8080", router)
}
