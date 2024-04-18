package departments

import (
	"html/template"
	"indodev-test/entities"
	"indodev-test/models/departmentmodel"
	"net/http"
)

func Get(w http.ResponseWriter, r *http.Request) {
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

func Add(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/departments/addform.html")
		if err != nil {
			panic(err)
		}
		temp.Execute(w, nil)
	}

	if r.Method == "POST" {
		var department entities.Departments

		department.DepartmentId = r.FormValue("id")
		department.DepartmentName = r.FormValue("name")

		if ok := departmentmodel.Insert(department); !ok {
			temp, _ := template.ParseFiles("views/departments/addform.html")
			temp.Execute(w, nil)
		}

		http.Redirect(w, r, "/departments", http.StatusSeeOther)
	}
}

func Edit(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/departments/editform.html")
		if err != nil {
			panic(err)
		}

		id := r.URL.Query().Get("id")

		department := departmentmodel.Detail(id)
		data := map[string]any{
			"department": department,
		}

		temp.Execute(w, data)
	}

	if r.Method == "POST" {
		var department entities.Departments

		id := r.FormValue("id")
		department.DepartmentName = r.FormValue("name")

		if ok := departmentmodel.Update(id, department); !ok {
			http.Redirect(w, r, r.Header.Get("Referer"), http.StatusSeeOther)
			return
		}

		http.Redirect(w, r, "/departments", http.StatusSeeOther)
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	if err := departmentmodel.Delete(id); err != nil {
		panic(err)
	}

	http.Redirect(w, r, "/departments", http.StatusSeeOther)
}
