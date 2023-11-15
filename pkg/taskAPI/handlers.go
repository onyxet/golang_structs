package taskAPI

import (
	"encoding/json"
	"net/http"
)

func tasks(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	dateParam := queryParams.Get("date")

	var filteredTasks []Task
	if dateParam != "" {
		for _, task := range FakeTasks {
			if task.Date == dateParam {
				filteredTasks = append(filteredTasks, task)
			}
		}
	} else {
		filteredTasks = FakeTasks
	}
	resp, err := json.MarshalIndent(filteredTasks, "", "  ")
	if err != nil {
		http.Error(w, "Error marshalling JSON", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}
