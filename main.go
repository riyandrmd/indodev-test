package main

import (
	"indodev-test/config"
	"indodev-test/controllers/departments"
	"indodev-test/controllers/homecontroller"
	"log"
	"net/http"
)

func main() {
	config.ConnectDB()

	// 1.Homepage
	http.HandleFunc("/", homecontroller.Homepage)

	// 2.Department
	http.HandleFunc("/departments", departments.Get)
	http.HandleFunc("/departments/add", departments.Add)
	http.HandleFunc("/departments/edit", departments.Edit)
	http.HandleFunc("/departments/delete", departments.Delete)

	log.Println("Server running on port 8080")
	http.ListenAndServe(":8080", nil)
}
