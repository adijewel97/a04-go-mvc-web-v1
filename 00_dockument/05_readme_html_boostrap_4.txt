Baik! Kita akan memodifikasi backend dan frontend agar:
✅ Jika input kosong → tampilkan semua data bank.
✅ Jika input diisi → tampilkan satu bank berdasarkan KODE_ERP.
✅ Jika tidak ditemukan → tampilkan pesan "Data tidak ada".

📌 1. Perbaiki Model models/bank.go
Sekarang query akan menyesuaikan apakah KODE_ERP kosong atau tidak.

go
Copy
Edit
package models

import (
	"a04-go-mvc-web-v1/config"
	"fmt"
)

type Bank struct {
	KODE_ERP  string `json:"kode_erp"`
	KODE_BANK string `json:"kode_bank"`
	NAMA_BANK string `json:"nama_bank"`
}

func GetMstBank(kodeERP string) ([]Bank, error) {
	var rows any
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
		return nil, fmt.Errorf("Data tidak ditemukan")
	}

	return banks, nil
}
📌 Perubahan:
✅ Jika kodeERP kosong → ambil semua data.
✅ Jika kodeERP diisi → ambil berdasarkan KODE_ERP.
✅ Jika tidak ditemukan, return error "Data tidak ditemukan".

📌 2. Perbaiki Controller controllers/bank_controller.go
Tambahkan logika untuk menangani data tidak ditemukan.

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

	banks, err := models.GetMstBank(kodeERP)
	if err != nil {
		http.Error(w, `{"message": "Data tidak ditemukan"}`, http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(banks)
}
📌 Perubahan:
✅ Tangani error "Data tidak ditemukan" → kirim { "message": "Data tidak ditemukan" } ke frontend.

📌 3. Perbaiki master_bank.html
Tambahkan logika untuk menangani pesan "Data tidak ada".

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
            <input type="text" class="form-control" id="kodeErp" placeholder="Kosongkan untuk melihat semua data">
        </div>
        <button class="btn btn-primary" onclick="fetchBanks()">Cari</button>

        <div id="message" class="alert alert-warning mt-3 d-none"></div>

        <table class="table table-bordered mt-4 d-none" id="bank-table-container">
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
            let url = "http://localhost:8080/banks"; // Default semua data
            if (kodeErp) {
                url += `?kode_erp=${kodeErp}`;
            }

            fetch(url)
                .then(response => {
                    if (!response.ok) {
                        return response.json().then(err => { throw err; });
                    }
                    return response.json();
                })
                .then(data => {
                    let tableContainer = document.getElementById("bank-table-container");
                    let message = document.getElementById("message");
                    let table = document.getElementById("bank-table");

                    // Reset tampilan
                    table.innerHTML = "";
                    message.classList.add("d-none");
                    tableContainer.classList.remove("d-none");

                    data.forEach(bank => {
                        let row = `<tr>
                            <td>${bank.kode_erp}</td>
                            <td>${bank.kode_bank}</td>
                            <td>${bank.nama_bank}</td>
                        </tr>`;
                        table.innerHTML += row;
                    });
                })
                .catch(error => {
                    document.getElementById("message").innerHTML = error.message;
                    document.getElementById("message").classList.remove("d-none");
                    document.getElementById("bank-table-container").classList.add("d-none");
                });
        }
    </script>

    <script src="https://cdn.jsdelivr.net/npm/bootstrap@5.3.0/dist/js/bootstrap.bundle.min.js"></script>
</body>
</html>
📌 Perubahan di frontend:
✅ Jika KODE_ERP kosong → tampilkan semua bank.
✅ Jika KODE_ERP ditemukan → tampilkan satu bank.
✅ Jika tidak ditemukan → tampilkan pesan "Data tidak ada".
✅ Gunakan Bootstrap alert untuk menampilkan pesan.

📌 4. Jalankan Server
Jalankan proyek dengan:

sh
Copy
Edit
go run main.go
Kemudian, buka di browser:
🔹 Master Bank: http://localhost:8080/master_bank.html

✅ Hasil Akhir
🔹 User bisa mengosongkan input untuk melihat semua bank.
🔹 Jika input diisi dengan KODE_ERP yang valid, muncul satu data.
🔹 Jika KODE_ERP tidak ditemukan, muncul pesan "Data tidak ada".
🔹 Bootstrap digunakan untuk tampilan yang rapi dan responsif.

🚀 Sekarang halaman master_bank.html sudah lebih fleksibel! 🚀

--- perbaikan
rows.Close undefined (type any has no field or method Close)
ChatGPT said:
Sepertinya terjadi masalah terkait penggunaan tipe data any untuk variabel rows. Pada Go, rows yang diterima dari fungsi config.DB.Query() harus dideklarasikan sebagai tipe yang sesuai dengan objek yang dihasilkan, dalam hal ini adalah tipe *sql.Rows. any adalah tipe yang terlalu umum dan tidak memiliki metode Close().

Untuk memperbaikinya, kita perlu mengganti any dengan tipe yang tepat, yaitu *sql.Rows. Berikut adalah perbaikan pada kode model models/bank.go:

Perbaiki Model models/bank.go
go
Copy
Edit
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
		return nil, fmt.Errorf("Data tidak ditemukan")
	}

	return banks, nil
}
Penjelasan Perubahan:
Gantilah deklarasi var rows any menjadi var rows *sql.Rows. Ini memastikan bahwa rows memiliki tipe yang benar dan memiliki metode Close().
*sql.Rows adalah tipe yang diterima ketika menggunakan config.DB.Query() yang mengembalikan hasil query dalam bentuk baris-baris yang bisa diakses dengan Next(), Scan(), dan ditutup dengan Close().
Dengan perubahan ini, seharusnya error terkait rows.Close tidak muncul lagi. Coba jalankan kembali aplikasi Anda setelah melakukan perubahan ini.

error strings should not be capitalized (ST1005)
ChatGPT said:
Error strings should not be capitalized (ST1005) biasanya muncul dari golangci-lint atau golangci-lint linting rules yang diaktifkan di proyek Anda. Ini mengindikasikan bahwa di dalam kode Anda ada string yang tidak mengikuti konvensi Go tentang penulisan string error.

Di Go, konvensi penulisan pesan error adalah untuk menulis pesan error dengan huruf kecil di awal, kecuali jika itu merupakan nama properti atau nama khusus.

Perbaikan Pesan Error di models/bank.go
Pada kode Anda, pada bagian ini:

go
Copy
Edit
return nil, fmt.Errorf("Data tidak ditemukan")
Pesan error "Data tidak ditemukan" dimulai dengan huruf kapital, yang bertentangan dengan konvensi Go. Anda harus menuliskannya dengan huruf kecil di awal, seperti berikut:

go
Copy
Edit
return nil, fmt.Errorf("data tidak ditemukan")
Perbaikan Kode Lengkap:
go
Copy
Edit
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
Penjelasan Perubahan:
Mengganti "Data tidak ditemukan" menjadi "data tidak ditemukan" agar sesuai dengan konvensi penulisan pesan error di Go yang menggunakan huruf kecil di awal.
Coba jalankan kembali proyek Anda setelah melakukan perubahan ini, dan linting error ST1005 seharusnya hilang.