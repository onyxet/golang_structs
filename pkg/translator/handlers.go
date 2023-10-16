package translator

import (
	"log"
	"net/http"
)

func translateHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form data", http.StatusBadRequest)
		return
	}
	client, err := CreateTranslateClient()
	if err != nil {
		log.Fatalf("Error creating Translate API client: %v", err)
	}
	sourceLang := r.FormValue("sourceLang")
	targetLang := r.FormValue("targetLang")
	text := r.FormValue("text")
	if sourceLang == "" || targetLang == "" || text == "" {
		http.Error(w, "Missing parameters", http.StatusBadRequest)
		return
	}
	translatedText, err := TranslateText(client, sourceLang, targetLang, text)
	if err != nil {
		http.Error(w, "Error translating text", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(translatedText))
}
