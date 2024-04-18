package main

import (
	"indodev-test/config"
	"indodev-test/controllers/homecontroller"
	"log"
	"net/http"
)

func main() {
	config.ConnectDB()

	http.HandleFunc("/", homecontroller.Homepage)

	log.Println("Server running on port 8080")
	http.ListenAndServe(":8080", nil)
}
