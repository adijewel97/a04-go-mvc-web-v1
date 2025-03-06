package routes

import (
	"a04-go-mvc-web-v1/controllers"
	"net/http"

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
