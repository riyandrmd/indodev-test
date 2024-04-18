package departments

import (
	"html/template"
	"indodev-test/models/departmentmodel"
	"net/http"
)

func GetAll(w http.ResponseWriter, r *http.Request) {
	departments := departmentmodel.GetAll()
	data := map[string]any{
		"departments": departments,
	}

	temp, err := template.ParseFiles("views/departments/index.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(w, data)
}
