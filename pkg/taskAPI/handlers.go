package taskAPI

import (
	"encoding/json"
	"net/http"
)

func tasks(w http.ResponseWriter, r *http.Request) {
	resp, err := json.MarshalIndent(FakeTasks, "", "  ")
	if err != nil {
		http.Error(w, "Error marshalling JSON", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}
