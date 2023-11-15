package taskAPI

import (
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/tasks", tasks).Methods("GET")
	return r
}
