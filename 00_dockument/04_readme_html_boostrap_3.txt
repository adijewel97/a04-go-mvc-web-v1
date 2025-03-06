0020400' adalah input an dari file master_bank.html
ChatGPT said:
Baik, jika KODE_ERP berasal dari input di halaman master_bank.html, kita harus mengubah beberapa hal:

Menambahkan input form di master_bank.html untuk memasukkan KODE_ERP.
Mengubah route dan controller agar bisa menerima KODE_ERP sebagai parameter dari frontend.
Memperbaiki model GetMstBank agar menerima KODE_ERP sebagai parameter.
📌 1. Perbaiki Model models/bank.go
Sekarang GetMstBank menerima kodeERP sebagai parameter.

go
Copy
Edit
package models

import "a04-go-mvc-web-v1/config"

type Bank struct {
	KODE_ERP  string `json:"kode_erp"`
	KODE_BANK string `json:"kode_bank"`
	NAMA_BANK string `json:"nama_bank"`
}

func GetMstBank(kodeERP string) ([]Bank, error) {
	rows, err := config.DB.Query("SELECT KODE_ERP, KODE_BANK, NAMA_BANK FROM MASTER_BANK WHERE KODE_ERP = :1", kodeERP)
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

	return banks, nil
}
📌 Perubahan:
✅ kodeERP dikirim dari frontend dan digunakan dalam query.
✅ Gunakan :1 (bind variable) untuk keamanan SQL Injection.

📌 2. Perbaiki Controller controllers/bank_controller.go
Sekarang kita menerima kode_erp dari URL query parameter (?kode_erp=0020400).

go
Copy
Edit
package controllers

import (
	"a04-go-mvc-web-v1/models"
	"encoding/json"
	"net/http"
)

func GetBanks(w http.ResponseWriter, r *http.Request) {
	kodeERP := r.URL.Query().Get("kode_erp") // Ambil kode_erp dari query parameter
	if kodeERP == "" {
		http.Error(w, "Kode ERP tidak boleh kosong", http.StatusBadRequest)
		return
	}

	banks, err := models.GetMstBank(kodeERP)
	if err != nil {
		http.Error(w, "Gagal mendapatkan data bank", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(banks)
}
📌 Perubahan:
✅ Ambil kode_erp dari query parameter.
✅ Jika kosong, kembalikan error 400.

📌 3. Perbaiki Route di routes/routes.go
Tambahkan endpoint yang menerima query parameter kode_erp.

go
Copy
Edit
package routes

import (
	"a04-go-mvc-web-v1/controllers"
	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	r := mux.NewRouter()

	// API Users
	r.HandleFunc("/users", controllers.GetUsers).Methods("GET")

	// API Banks (dengan parameter kode_erp)
	r.HandleFunc("/banks", controllers.GetBanks).Methods("GET")

	// Serve static files
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./public/")))

	return r
}
📌 Endpoint baru:
🔹 http://localhost:8080/banks?kode_erp=0020400

📌 4. Perbaiki master_bank.html
Sekarang kita tambahkan input untuk KODE_ERP, dan tombol pencarian.

html
Copy
Edit
<!DOCTYPE html>
<html lang="id">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Master Bank</title>
    <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.0/dist/css/bootstrap.min.css" rel="stylesheet">
</head>
<body>
    <nav class="navbar navbar-expand-lg navbar-dark bg-dark">
        <div class="container">
            <a class="navbar-brand" href="index.html">MyApp</a>
            <button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarNav">
                <span class="navbar-toggler-icon"></span>
            </button>
            <div class="collapse navbar-collapse" id="navbarNav">
                <ul class="navbar-nav">
                    <li class="nav-item">
                        <a class="nav-link" href="user.html">User</a>
                    </li>
                    <li class="nav-item">
                        <a class="nav-link active" href="master_bank.html">Master Bank</a>
                    </li>
                </ul>
            </div>
        </div>
    </nav>

    <div class="container mt-4">
        <h2>Data Master Bank</h2>

        <!-- Form Input KODE_ERP -->
        <div class="mb-3">
            <label for="kodeErp" class="form-label">Masukkan Kode ERP:</label>
            <input type="text" class="form-control" id="kodeErp" placeholder="Contoh: 0020400">
        </div>
        <button class="btn btn-primary" onclick="fetchBanks()">Cari</button>

        <table class="table table-bordered mt-4">
            <thead>
                <tr>
                    <th>Kode ERP</th>
                    <th>Kode Bank</th>
                    <th>Nama Bank</th>
                </tr>
            </thead>
            <tbody id="bank-table"></tbody>
        </table>
    </div>

    <script>
        function fetchBanks() {
            let kodeErp = document.getElementById("kodeErp").value;
            if (!kodeErp) {
                alert("Kode ERP tidak boleh kosong");
                return;
            }

            fetch(`http://localhost:8080/banks?kode_erp=${kodeErp}`)
                .then(response => response.json())
                .then(data => {
                    let table = document.getElementById("bank-table");
                    table.innerHTML = "";
                    data.forEach(bank => {
                        let row = `<tr>
                            <td>${bank.kode_erp}</td>
                            <td>${bank.kode_bank}</td>
                            <td>${bank.nama_bank}</td>
                        </tr>`;
                        table.innerHTML += row;
                    });
                })
                .catch(error => console.error("Error fetching data:", error));
        }
    </script>

    <script src="https://cdn.jsdelivr.net/npm/bootstrap@5.3.0/dist/js/bootstrap.bundle.min.js"></script>
</body>
</html>
📌 Fitur di master_bank.html:
✅ Input KODE_ERP untuk filter data bank.
✅ Tombol "Cari" untuk memuat data berdasarkan KODE_ERP.
✅ Menampilkan hasil pencarian di tabel Bootstrap.

📌 5. Jalankan Server
Buka terminal di folder proyek, lalu jalankan:

sh
Copy
Edit
go run main.go
Kemudian, buka di browser:
🔹 Master Bank: http://localhost:8080/master_bank.html

✅ Hasil Akhir
🔹 User bisa memasukkan KODE_ERP di form input.
🔹 Tombol "Cari" akan menampilkan data bank sesuai KODE_ERP yang dimasukkan.
🔹 Menggunakan Bootstrap agar tampilan lebih bagus.
🔹 Backend dan frontend sudah saling terhubung.

🚀 Sekarang master_bank.html bisa menampilkan data sesuai input dari user! 🚀