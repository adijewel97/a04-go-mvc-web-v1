Berikut adalah struktur dan isi file yang diperlukan untuk membuat dua menu: User dan Master Bank, menggunakan Bootstrap.

📂 Struktur Folder
go
Copy
Edit
a04-go-mvc-web-v1/
│── main.go
│── config/
│── controllers/
│── models/
│── routes/
│── public/
│   ├── index.html
│   ├── user.html
│   ├── master_bank.html
│   ├── css/
│   │   ├── style.css
│   ├── js/
│   │   ├── script.js
│── go.mod
│── go.sum
1. index.html (Halaman Utama)
File ini akan memiliki dua menu: User dan Master Bank.

html
Copy
Edit
<!DOCTYPE html>
<html lang="id">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Dashboard</title>
    <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.0/dist/css/bootstrap.min.css" rel="stylesheet">
</head>
<body>
    <nav class="navbar navbar-expand-lg navbar-dark bg-dark">
        <div class="container">
            <a class="navbar-brand" href="#">MyApp</a>
            <button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarNav">
                <span class="navbar-toggler-icon"></span>
            </button>
            <div class="collapse navbar-collapse" id="navbarNav">
                <ul class="navbar-nav">
                    <li class="nav-item">
                        <a class="nav-link" href="user.html">User</a>
                    </li>
                    <li class="nav-item">
                        <a class="nav-link" href="master_bank.html">Master Bank</a>
                    </li>
                </ul>
            </div>
        </div>
    </nav>

    <div class="container mt-4">
        <h1>Selamat Datang di Dashboard</h1>
        <p>Pilih menu di atas untuk melihat data.</p>
    </div>

    <script src="https://cdn.jsdelivr.net/npm/bootstrap@5.3.0/dist/js/bootstrap.bundle.min.js"></script>
</body>
</html>
2. user.html (Halaman User)
html
Copy
Edit
<!DOCTYPE html>
<html lang="id">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Data User</title>
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
                        <a class="nav-link active" href="user.html">User</a>
                    </li>
                    <li class="nav-item">
                        <a class="nav-link" href="master_bank.html">Master Bank</a>
                    </li>
                </ul>
            </div>
        </div>
    </nav>

    <div class="container mt-4">
        <h2>Daftar Pengguna</h2>
        <table class="table table-bordered">
            <thead>
                <tr>
                    <th>ID</th>
                    <th>Nama</th>
                    <th>Username</th>
                </tr>
            </thead>
            <tbody id="user-table"></tbody>
        </table>
    </div>

    <script>
        fetch("http://localhost:8080/users.xml")
        .then(response => response.text())
        .then(str => (new window.DOMParser()).parseFromString(str, "text/xml"))
        .then(data => {
            let users = data.getElementsByTagName("user");
            let table = document.getElementById("user-table");

            for (let user of users) {
                let id = user.getElementsByTagName("id")[0].textContent;
                let name = user.getElementsByTagName("name")[0].textContent;
                let username = user.getElementsByTagName("username")[0].textContent;

                let row = `<tr>
                            <td>${id}</td>
                            <td>${name}</td>
                            <td>${username}</td>
                        </tr>`;
                table.innerHTML += row;
            }
        });
    </script>

    <script src="https://cdn.jsdelivr.net/npm/bootstrap@5.3.0/dist/js/bootstrap.bundle.min.js"></script>
</body>
</html>
3. master_bank.html (Halaman Master Bank)
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
        <h2>Daftar Bank</h2>
        <table class="table table-bordered">
            <thead>
                <tr>
                    <th>ID Bank</th>
                    <th>Nama Bank</th>
                </tr>
            </thead>
            <tbody id="bank-table"></tbody>
        </table>
    </div>

    <script>
        fetch("http://localhost:8080/banks.xml")
        .then(response => response.text())
        .then(str => (new window.DOMParser()).parseFromString(str, "text/xml"))
        .then(data => {
            let banks = data.getElementsByTagName("bank");
            let table = document.getElementById("bank-table");

            for (let bank of banks) {
                let id = bank.getElementsByTagName("id")[0].textContent;
                let name = bank.getElementsByTagName("name")[0].textContent;

                let row = `<tr>
                            <td>${id}</td>
                            <td>${name}</td>
                        </tr>`;
                table.innerHTML += row;
            }
        });
    </script>

    <script src="https://cdn.jsdelivr.net/npm/bootstrap@5.3.0/dist/js/bootstrap.bundle.min.js"></script>
</body>
</html>
📌 Penjelasan
✅ index.html → Halaman utama dengan 2 menu
✅ user.html → Menampilkan data User dalam tabel dari API /users.xml
✅ master_bank.html → Menampilkan data Master Bank dalam tabel dari API /banks.xml
✅ Bootstrap digunakan untuk desain yang responsif dan menarik
✅ AJAX (fetch API) untuk mengambil data dari backend

Coba jalankan proyekmu dan akses halaman index.html! 🚀