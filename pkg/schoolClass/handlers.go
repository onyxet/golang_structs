package schoolClass

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func students(w http.ResponseWriter, r *http.Request) {
	fakeStudents := GenerateStudents()
	resp, err := json.MarshalIndent(fakeStudents, "", "  ")
	if err != nil {
		http.Error(w, "Error marshalling JSON", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func studentByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	studentID := vars["id"]

	id, err := strconv.Atoi(studentID)
	if err != nil {
		http.Error(w, "Invalid student ID", http.StatusBadRequest)
		return
	}

	student, err := GetStudentByID(id)
	if err != nil {
		http.Error(w, "Student not found", http.StatusNotFound)
		return
	}

	resp, err := json.MarshalIndent(student, "", "  ")
	if err != nil {
		http.Error(w, "Error marshalling JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}
