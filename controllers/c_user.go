package controllers

import (
	"a04-go-mvc-web-v1/models"
	"encoding/json"
	"encoding/xml"
	"net/http"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := models.GetUsers()
	if err != nil {
		http.Error(w, "Gagal mendapatkan data user", http.StatusInternalServerError)
		return
	}

	// Cek format response berdasarkan query parameter (?format=xml/json)
	format := r.URL.Query().Get("format")
	if format == "xml" {
		w.Header().Set("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(users)
	} else {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(users)
	}
}
