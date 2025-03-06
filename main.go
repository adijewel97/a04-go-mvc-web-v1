package main

import (
	"a04-go-mvc-web-v1/config"
	"a04-go-mvc-web-v1/routes"
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
