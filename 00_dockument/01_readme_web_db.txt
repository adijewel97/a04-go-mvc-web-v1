Berikut adalah contoh sederhana implementasi pola MVC di Golang untuk koneksi ke database Oracle 12c menggunakan godror.

Struktur Proyek
go
Copy
Edit
a03-my-go-project/
│── main.go
│── config/
│   ├── database.go
│   ├── config.properties
│── models/
│   ├── user.go
│── controllers/
│   ├── user_controller.go
│── routes/
│   ├── routes.go
│── go.mod
│── go.sum
1. Konfigurasi Database (config/database.go)
go
Copy
Edit
package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/godror/godror"
)

var DB *sql.DB

func ConnectDB() {
	dsn := "USERADISPPOBNTL/adis123@127.0.0.1:1521/adis.iconpln.co.id"
	db, err := sql.Open("godror", dsn)
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal("Database not reachable: ", err)
	}

	fmt.Println("Connected to Oracle Database!")
	DB = db
}
2. Model (models/user.go)
go
Copy
Edit
package models

import (
	"a03-my-go-project/config"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func GetUsers() ([]User, error) {
	rows, err := config.DB.Query("SELECT ID, NAME FROM USERS")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Name); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}
3. Controller (controllers/user_controller.go)
go
Copy
Edit
package controllers

import (
	"a03-my-go-project/models"
	"encoding/json"
	"net/http"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := models.GetUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
4. Routing (routes/routes.go)
go
Copy
Edit
package routes

import (
	"a03-my-go-project/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/users", controllers.GetUsers).Methods("GET")
	return r
}
5. Entry Point (main.go)
go
Copy
Edit
package main

import (
	"a03-my-go-project/config"
	"a03-my-go-project/routes"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.ConnectDB()

	r := routes.SetupRoutes()
	fmt.Println("Server running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
6. Konfigurasi (config/config.properties)
ini
Copy
Edit
db.url=127.0.0.1:1521/adis.iconpln.co.id
db.username=USERADISPPOBNTL
db.password=adis123
7. Install Dependensi
Jalankan perintah berikut untuk menginstal dependensi yang dibutuhkan:

sh
Copy
Edit
go mod init a03-my-go-project
go get github.com/godror/godror
go get github.com/gorilla/mux
8. Menjalankan Aplikasi
Jalankan server dengan perintah:

sh
Copy
Edit
go run main.go
Akses data pengguna dengan membuka http://localhost:8080/users di browser atau menggunakan curl:

sh
Copy
Edit
curl http://localhost:8080/users
Dengan struktur ini, proyek sudah menerapkan pola MVC dengan Golang, Godror, dan Oracle 12c. 🚀
