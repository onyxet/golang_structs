package weather

import (
	"encoding/json"
	"net/http"
)

func weatherHandler(w http.ResponseWriter, r *http.Request) {
	city := r.URL.Query().Get("city")
	weatherData, err := GetWeather(city)
	if err != nil {
		http.Error(w, "Error getting weather data", http.StatusInternalServerError)
		return
	}
	jsonData, err := json.Marshal(weatherData)
	if err != nil {
		http.Error(w, "Error marshalling JSON", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}
