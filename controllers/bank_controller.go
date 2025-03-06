package controllers

import (
	"a04-go-mvc-web-v1/models"
	"encoding/json"
	"net/http"
)

func GetBanks(w http.ResponseWriter, r *http.Request) {
	kodeERP := r.URL.Query().Get("kode_erp") // Ambil kode_erp dari query parameter

	banks, err := models.GetMstBank(kodeERP)
	if err != nil {
		http.Error(w, `{"message": "Data tidak ditemukan"}`, http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(banks)
}
