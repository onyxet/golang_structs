package weather

import (
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/weather", weatherHandler).Methods("GET")
	return r
}
