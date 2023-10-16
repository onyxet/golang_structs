package schoolClass

import (
	"github.com/gorilla/mux"
	"net/http"
)

func Router() *mux.Router {
	r := mux.NewRouter()
	r.Handle("/students", Authenticator(http.HandlerFunc(students))).Methods("GET")
	r.Handle("/students/{id:[0-9]+}", Authenticator(http.HandlerFunc(studentByID))).Methods("GET")

	return r
}
