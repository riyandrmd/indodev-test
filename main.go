package main

import (
	"indodev-test/config"
	"indodev-test/controllers/departments"
	"indodev-test/controllers/divisioncontroller"
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

	// 3.Divions
	http.HandleFunc("/divisions", divisioncontroller.Get)
	http.HandleFunc("/divisions/add", divisioncontroller.Add)
	http.HandleFunc("/divisions/edit", divisioncontroller.Edit)
	http.HandleFunc("/divisions/delete", divisioncontroller.Delete)

	log.Println("Server running on port 8080")
	http.ListenAndServe(":8080", nil)
}
