package translator

import "github.com/gorilla/mux"

func Router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/translate", translateHandler).Methods("POST")
	return r
}
