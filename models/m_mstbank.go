package models

import (
	"a04-go-mvc-web-v1/config"
	"database/sql"
	"fmt"
)

type Bank struct {
	KODE_ERP  string `json:"kode_erp"`
	KODE_BANK string `json:"kode_bank"`
	NAMA_BANK string `json:"nama_bank"`
}

func GetMstBank(kodeERP string) ([]Bank, error) {
	var rows *sql.Rows // Gantilah 'any' dengan '*sql.Rows'
	var err error

	// Jika kodeERP kosong, ambil semua data
	if kodeERP == "" {
		rows, err = config.DB.Query("SELECT KODE_ERP, KODE_BANK, NAMA_BANK FROM MASTER_BANK")
	} else {
		rows, err = config.DB.Query("SELECT KODE_ERP, KODE_BANK, NAMA_BANK FROM MASTER_BANK WHERE KODE_ERP = :1", kodeERP)
	}

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var banks []Bank
	for rows.Next() {
		var bank Bank
		if err := rows.Scan(&bank.KODE_ERP, &bank.KODE_BANK, &bank.NAMA_BANK); err != nil {
			return nil, err
		}
		banks = append(banks, bank)
	}

	// Jika tidak ada data ditemukan, return slice kosong
	if len(banks) == 0 {
		return nil, fmt.Errorf("data tidak ditemukan") // Ganti dengan huruf kecil
	}

	return banks, nil
}
